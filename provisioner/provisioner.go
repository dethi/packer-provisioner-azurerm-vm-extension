package provisioner

import (
	"context"
	"fmt"
	"os"

	"github.com/dethi/packer-provisioner-azurerm-vm-extension/azclient"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/packer/helper/config"
	"github.com/hashicorp/packer/packer"
)

type Provisioner struct {
	config Config

	client compute.VirtualMachineExtensionsClient
	ctx    context.Context
	cancel context.CancelFunc
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := config.Decode(&p.config, &config.DecodeOpts{
		Interpolate:        true,
		InterpolateContext: &p.config.ctx,
	}, raws...)
	if err != nil {
		return err
	}

	clientId := p.config.ClientId
	clientSecret := p.config.ClientSecret
	tenantId := p.config.TenantId
	subId := p.config.SubscriptionId

	authorizer, err := azclient.GetAuthorizer(clientId, clientSecret, tenantId)
	if err != nil {
		return err
	}
	p.client = azclient.GetVMExtensionsClient(subId, authorizer)

	p.ctx, p.cancel = context.WithCancel(context.Background())
	return nil
}

func (p *Provisioner) Provision(ui packer.Ui, comm packer.Communicator) error {
	name := p.config.Name
	resGroup := p.config.ResourceGroup
	vmName := p.config.VmName

	ui.Say("Provisioning Virtual Machine extension...")

	var future azure.Future
	if p.config.DeleteExtension {
		if f, err := p.client.Delete(p.ctx, resGroup, vmName, name); err == nil {
			ui.Message("Removing extension...")
			future = f.Future
		} else {
			return err
		}
	} else {
		extension, err := newExtension(p.config)
		if err != nil {
			return err
		}

		if f, err := p.client.CreateOrUpdate(p.ctx, resGroup, vmName, name, *extension); err == nil {
			ui.Message("Adding extension...")
			future = f.Future
		} else {
			return err
		}
	}

	err := future.WaitForCompletionRef(p.ctx, p.client.Client)
	if err != nil {
		return err
	}

	ui.Message("--> done!")
	return nil
}

func (p *Provisioner) Cancel() {
	if p.cancel != nil {
		p.cancel()
	}

	os.Exit(0)
}

func newExtension(config Config) (*compute.VirtualMachineExtension, error) {
	extension := compute.VirtualMachineExtension{
		Location: &config.Location,
		VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
			Publisher:               &config.Publisher,
			Type:                    &config.ExtensionType,
			TypeHandlerVersion:      &config.TypeHandlerVersion,
			AutoUpgradeMinorVersion: &config.AutoUpgradeMinorVersion,
		},
	}

	if config.SettingsFile != "" {
		settings, err := readJsonFile(config.SettingsFile)
		if err != nil {
			return nil, fmt.Errorf("unable to read settings: %v", err)
		}
		extension.VirtualMachineExtensionProperties.Settings = &settings
	}

	if config.ProtectedSettingsFile != "" {
		protectedSettings, err := readJsonFile(config.ProtectedSettingsFile)
		if err != nil {
			return nil, fmt.Errorf("unable to read protected_settings: %v", err)
		}
		extension.VirtualMachineExtensionProperties.ProtectedSettings = &protectedSettings
	}

	return &extension, nil
}

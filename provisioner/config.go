package provisioner

import (
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	ctx                 interpolate.Context

	ClientId       string `mapstructure:"client_id"`
	ClientSecret   string `mapstructure:"client_secret"`
	TenantId       string `mapstructure:"tenant_id"`
	SubscriptionId string `mapstructure:"subscription_id"`

	ResourceGroup string `mapstructure:"resource_group_name"`
	VmName        string `mapstructure:"vm_name"`
	Location      string `mapstructure:"location"`

	Name                    string `mapstructure:"name"`
	Publisher               string `mapstructure:"publisher"`
	ExtensionType           string `mapstructure:"extension_type"`
	TypeHandlerVersion      string `mapstructure:"type_handler_version"`
	AutoUpgradeMinorVersion bool   `mapstructure:"auto_upgrade_minor_version"`
	DeleteExtension         bool   `mapstructure:"delete_extension"`

	SettingsFile          string `mapstructure:"settings_file"`
	ProtectedSettingsFile string `mapstructure:"protected_settings_file"`
}

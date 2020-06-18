package azclient

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-06-01/compute"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

const (
	USER_AGENT = "packer-provisioner-azurerm-vm-extension"
)

func GetVMExtensionsClient(subscriptionId string, authorizer autorest.Authorizer) compute.VirtualMachineExtensionsClient {
	extensionsClient := compute.NewVirtualMachineExtensionsClient(subscriptionId)
	extensionsClient.Authorizer = authorizer
	extensionsClient.UserAgent = USER_AGENT

	return extensionsClient
}

func GetAuthorizer(clientId, clientSecret, tenantId string) (autorest.Authorizer, error) {
	credentialsAuthorizer := auth.NewClientCredentialsConfig(clientId, clientSecret, tenantId)
	return credentialsAuthorizer.Authorizer()
}

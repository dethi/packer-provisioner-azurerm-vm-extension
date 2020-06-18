package main

import (
	"github.com/dethi/packer-provisioner-azurerm-vm-extension/provisioner"
	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(provisioner.Provisioner))
	server.Serve()
}

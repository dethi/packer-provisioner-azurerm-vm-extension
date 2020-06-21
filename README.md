# packer-provisioner-azurerm-vm-extension

`packer-provisioner-azurerm-vm-extension` is a provisioner plugin for Packer that let you add/remove Azure VM Extension during the provisioning phase.

**Tested with Packer <=1.3**. The plugin may need to be adapted for Packer >=1.4 because of the addition of HCL2.

## Use Cases

- Reuse existing provisioning stack (like DSC script)
- Resolves [packer#6548](https://github.com/hashicorp/packer/issues/6548) and [packer#8626](https://github.com/hashicorp/packer/issues/8626)

## Example

[An example](https://github.com/dethi/packer-provisioner-azurerm-vm-extension/tree/master/example) is provided for provisioning a Windows VM with the DSC Extension.

However, any Azure VM Extension should work.

## Build & Usage

```sh
# build binary
$ go build -v .

# move plugin to specific directory
$ mkdir -p $HOME/.packer.d/plugins
$ mv packer-provisioner-azurerm-vm-extension $HOME/.packer.d/plugins

# run packer
$ cd example/
$ packer build -var 'foo=bar' [...] ./packer_example.json
```

For more information, read the [Packer plugin documentation](https://www.packer.io/docs/extending/plugins).

package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sam-myers/terraform-provider-linux/linux"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: linux.Provider})
}

package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/sam-myers/terraform-provider-remote/template"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: template.Provider})
}

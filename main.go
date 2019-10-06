package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-scaffolding/scaffolding" // change this to the import path of your provider
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: scaffolding.Provider})
}

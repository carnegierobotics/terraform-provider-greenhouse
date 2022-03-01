package main

import (
	"github.com/carnegierobotics/terraform-provider-greenhouse/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: greenhouse.Provider,
	})
}

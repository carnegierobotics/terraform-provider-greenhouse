package main

import (
	"github.com/carnegierobotics/terraform-provider-greenhouse/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: greenhouse.Provider,
	})
}

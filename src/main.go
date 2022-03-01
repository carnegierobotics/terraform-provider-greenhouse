package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/plugin"
  "github.com/carnegierobotics/terraform-provider-greenhouse/greenhouse"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: greenhouse.Provider,
  })
}

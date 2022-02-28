package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/plugin"
  "github.com/carnegierobotics/terraform-provider-greenhouse/internal/provider"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: provider.Provider
  })
}

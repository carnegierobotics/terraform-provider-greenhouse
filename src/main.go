package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/plugin"
  "github.com/carnegierobotics/greenhouse-client-go/internal/greenhouse"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: greenhouse.Provider
  })
}

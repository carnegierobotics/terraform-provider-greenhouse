package greenhouse

import (
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
  p := &schema.Provider {
    Schema: map[string]*schema.Schema{
      "token": {
        Type:        schema.TypeString,
        Optional:    false,
        DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_TOKEN", nil),
        Description: "This is an API token for Greenhouse.",
      },
      "url": {
        Type:        schema.TypeString,
        Optional:    false,
        DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_TOKEN", nil),
        Description: "The URL for Greenhouse.",
      },
    },
    ResourcesMap: map[string]*schema.Resource {
      "greenhouse_job" : resourceGreenhouseJob(),
      "greenhouse_user": resourceGreenhouseUser(),
    },
    DataSourcesMap: map[string]*schema.Resource {
      "greenhouse_jobs" : dataSourceGreenhouseJobs(),
      "greenhouse_users": dataSourceGreenhouseUsers(),
    },
  }
  p.ConfigureFunc = providerConfigure(p)
  return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
  return func(d *schema.ResourceData) (interface{}, error) {
    c, err = NewConfig(d)
    if err != nil {
      return nil, err
    }
    return c.Client()
  }
}

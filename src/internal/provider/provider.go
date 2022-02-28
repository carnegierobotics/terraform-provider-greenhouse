package greenhouse

import (
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
  p := &schema.Provider {
    Schema: map[string]*schema.Schema{
      "on_behalf_of": {
        Type:        schema.TypeString,
        Optional:    false,
        DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_ON_BEHALF_OF", nil),
        Description: "This is the user on whose behalf all actions will be audited.",
      },
      "token": {
        Type:        schema.TypeString,
        Optional:    false,
        DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_TOKEN", nil),
        Description: "This is an API token for Greenhouse.",
      },
      "jobs_url": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_JOBS_URL", "https://boards-api.greenhouse.io"),
        Description: "The URL for Greenhouse Job Boards API.",
      },
      "harvest_url": {
        Type:        schema.TypeString,
        Optional:    true,
        DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_HARVEST_URL", "https://harvest.greenhouse.io"),
        Description: "The URL for Greenhouse's Harvest API."
      }
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

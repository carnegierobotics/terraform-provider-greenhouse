package greenhouse

import (
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"strconv"
)

func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"on_behalf_of": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_ON_BEHALF_OF", nil),
				Description: "This is the user on whose behalf all actions will be audited.",
			},
			"jobs_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_JOBS_TOKEN", nil),
				Description: "The token to use for the Greenhouse Jobs API.",
			},
			"harvest_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_HARVEST_TOKEN", nil),
				Description: "The token to use for the Greenhouse Harvest API.",
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
				Description: "The URL for Greenhouse's Harvest API.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"greenhouse_close_reason": resourceGreenhouseCloseReason(),
			"greenhouse_department":   resourceGreenhouseDepartment(),
			"greenhouse_office":       resourceGreenhouseOffice(),
			"greenhouse_user":         resourceGreenhouseUser(),
			"greenhouse_job":          resourceGreenhouseJob(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			// "greenhouse_departments": dataSourceGreenhouseDepartment(),
			// "greenhouse_jobs"       : dataSourceGreenhouseJobs(),
			// "greenhouse_users"      : dataSourceGreenhouseUsers(),
		},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		harvest_url := d.Get("harvest_url").(string)
		harvest_token := d.Get("harvest_token").(string)
		on_behalf_of, err := strconv.Atoi(d.Get("on_behalf_of").(string))
		if err != nil {
			return nil, err
		}
		client := greenhouse.Client{BaseUrl: harvest_url, Token: harvest_token, OnBehalfOf: on_behalf_of}
		err = client.BuildResty()
		if err != nil {
			return nil, err
		}
		return &client, nil
	}
}

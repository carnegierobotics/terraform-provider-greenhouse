package greenhouse

import (
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func Provider() *schema.Provider {
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
			"greenhouse_activity_feed": resourceGreenhouseActivityFeed(),
      //"greenhouse_application": resourceGreenhouseApplication(),
      //"greenhouse_approval": resourceGreenhouseApproval(),
			"greenhouse_candidate":     resourceGreenhouseCandidate(),
			"greenhouse_close_reason":  resourceGreenhouseCloseReason(),
      //"greenhouse_custom_field": resourceGreenhouseCustomField(),
			"greenhouse_department":    resourceGreenhouseDepartment(),
			"greenhouse_job":           resourceGreenhouseJob(),
			"greenhouse_office":        resourceGreenhouseOffice(),
			"greenhouse_user":          resourceGreenhouseUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
      "greenhouse_activity_feed": dataSourceGreenhouseActivityFeed(),
      //"greenhouse_application": dataSourceGreenhouseApplication(),
      //"greenhouse_candidate": dataSourceGreenhouseCandidate(),
      //"greenhouse_close_reason": dataSourceGreenhouseCloseReason(),
      //"greenhouse_close_reasons": dataSourceGreenhouseCloseReasons(),
      //"greenhouse_custom_field": dataSourceGreenhouseCustomField(),
      //"greenhouse_custom_fields": dataSourceGreenhouseCustomFields(),
      //"greenhouse_demographic_answer": dataSourceGreenhouseDemographicAnswer(),
      //"greenhouse_demographic_answers": dataSourceGreenhouseDemographicAnswers(),
      //"greenhouse_demographic_question": dataSourceGreenhouseDemographicQuestion(),
      //"greenhouse_demographic_questions": dataSourceGreenhouseDemographicQuestions(),
      "greenhouse_department": dataSourceGreenhouseDepartment(),
			"greenhouse_departments": dataSourceGreenhouseDepartments(),
      "greenhouse_education_degree": dataSourceGreenhouseEducationDegree(),
      "greenhouse_education_discipline": dataSourceGreenhouseEducationDiscipline(),
      "greenhouse_education_school": dataSourceGreenhouseEducationSchool(),
      "greenhouse_job": dataSourceGreenhouseJob(),
			"greenhouse_jobs"       : dataSourceGreenhouseJobs(),
      "greenhouse_office": dataSourceGreenhouseOffice(),
      "greenhouse_prospect_pool": dataSourceGreenhouseProspectPool(),
      "greenhouse_rejection_reason": dataSourceGreenhouseRejectionReason(),
      "greenhouse_source": dataSourceGreenhouseSource(),
      "greenhouse_sources": dataSourceGreenhouseSources(),
      "greenhouse_user": dataSourceGreenhouseUser(),
      "greenhouse_user_role": dataSourceGreenhouseUserRole(),
      "greenhouse_users": dataSourceGreenhouseUsers(),
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

package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmployment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"company_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"title": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

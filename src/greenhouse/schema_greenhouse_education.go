package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEducation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"degree": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"degree_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"discipline": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"discipline_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_month": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end_year": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"school_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_month": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"start_year": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

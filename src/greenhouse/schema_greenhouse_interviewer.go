package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterviewer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"response_status": {
			Type:     schema.TypeString,
			Required: true,
		},
		"scorecard_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

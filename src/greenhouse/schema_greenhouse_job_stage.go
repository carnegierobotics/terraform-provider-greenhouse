package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJobStage() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"interviews": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterview(),
			},
		},
		"job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

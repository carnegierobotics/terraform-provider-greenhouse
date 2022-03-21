package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJobPost() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"content": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"demographic_question_set_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"external": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"first_published_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"internal": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"internal_content": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"live": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"questions": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDemographicQuestion(),
			},
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

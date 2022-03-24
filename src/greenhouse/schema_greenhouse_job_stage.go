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
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterview(),
			},
		},
		"id": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"job_id", "name"},
		},
		"job_id": {
			Type:          schema.TypeInt,
			Optional:      true,
			ConflictsWith: []string{"id"},
			RequiredWith:  []string{"name"},
		},
		"name": {
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"id"},
			RequiredWith:  []string{"name"},
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

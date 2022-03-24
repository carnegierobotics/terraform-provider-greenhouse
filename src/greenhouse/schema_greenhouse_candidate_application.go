package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCandidateApplication() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attachments": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseAttachment(),
			},
		},
		"initial_stage_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"referrer": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
	}
}

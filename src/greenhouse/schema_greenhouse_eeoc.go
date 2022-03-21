package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEEOC() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"candidate_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"disability_status": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
		"gender": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
		"race": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
		"submitted_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"veteran_status": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
	}
}

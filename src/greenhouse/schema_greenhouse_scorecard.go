package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScorecard() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"applicationId": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"attributes": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseScorecardAttribute(),
			},
		},
		"candidate_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"interview": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"interview_step": {
			Type:     schema.TypeMap,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"interviewer": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUser(),
			},
		},
		"overall_recommendation": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"questions": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseScorecardQuestion(),
			},
		},
		"ratings": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		"submitted_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"submitted_by": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUser(),
			},
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

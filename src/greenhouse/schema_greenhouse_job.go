package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJob() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"job_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"job_post_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"requisition_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"notes": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"anywhere": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"teams_and_responsibilities": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"how_to_sell_this_job": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"confidential": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"opened_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"closed_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"is_template": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"copied_from_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"department_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"departments": {
			Type:     schema.TypeList,
			MinItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"office_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"offices": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffice(),
			},
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"keyed_custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"hiring_team_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"hiring_team": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseHiringMember(),
				},
			},
		},
		"number_of_openings": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"template_job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"openings": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseJobOpening(),
			},
		},
		"opening_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	}
}

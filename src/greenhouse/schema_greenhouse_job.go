package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
			Type:     schema.TypeSet,
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
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffice(),
			},
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseCustomField(),
			},
		},
		"hiring_team_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"hiring_team": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseHiringTeam(),
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
		"opening_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseJobOpening(),
			},
		},
		"openings": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	}
}

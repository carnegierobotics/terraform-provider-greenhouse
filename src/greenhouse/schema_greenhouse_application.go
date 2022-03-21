package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApplication() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"answers": {
			Type:        schema.TypeSet,
			Description: "",
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseAnswer(),
			},
		},
		"applied_at": {
			Type:        schema.TypeString,
			Description: "The date of the application.",
			Computed:    true,
		},
		"attachments": {
			Type:        schema.TypeSet,
			Description: "",
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseAttachment(),
			},
		},
		"candidate_id": {
			Type:        schema.TypeInt,
			Description: "The ID of the candidate applying for this job.",
			Computed:    true,
		},
		"credited_to": {
			Type:        schema.TypeSet,
			Description: "The user who will receive credit for this application.",
			MaxItems:    1,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
		"current_stage": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"jobs": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseJob(),
			},
		},
		"job_post_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"keyed_custom_fields": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseKeyedCustomField(),
			},
		},
		"last_activity_at": {
			Type:        schema.TypeString,
			Description: "The date of the application's last activity.",
			Computed:    true,
		},
		"location": {
			Type:        schema.TypeList,
			Description: "The contents of a location question on a job post.",
			MaxItems:    1,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseLocation(),
			},
		},
		"prospect": {
			Type:        schema.TypeBool,
			Description: "The candidate is a prospect and has not yet applied.",
			Computed:    true,
		},
		"prospect_detail": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseProspectDetail(),
			},
		},
		"prospect_pool_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"prospect_pool_stage_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		// I suspect this is a typo in their docs, but just in case.
		"prospect_stage_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"prospective_department": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"prospective_office": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffice(),
			},
		},
		"referrer": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"rejected_at": {
			Type:        schema.TypeString,
			Description: "The date of the application's rejection.",
			Computed:    true,
		},
		"source": {
			Type:        schema.TypeSet,
			Description: "",
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseSource(),
			},
		},
		"source_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

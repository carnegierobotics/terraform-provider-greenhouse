package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
		"current_stage": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
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
		"from_stage_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"hire": {
			Type:     schema.TypeBool,
			Optional: true,
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
			Optional: true,
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
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseLocation(),
			},
		},
		"new_job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"new_stage_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"prospect": {
			Type:        schema.TypeBool,
			Description: "The candidate is a prospect and has not yet applied.",
			Computed:    true,
		},
		"prospect_detail": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
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
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"prospective_office": {
			Type:     schema.TypeSet,
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
		"reject": {
			Type:     schema.TypeBool,
			Optional: true,
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

func convertToApplicationList(list []interface{}) *[]greenhouse.Application {
	newList := make([]greenhouse.Application, len(list))
	for i := range list {
		newList[i] = list[i].(greenhouse.Application)
	}
	return &newList
}

func flattenApplications(ctx context.Context, list *[]greenhouse.Application) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApplication(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApplication(ctx context.Context, item *greenhouse.Application) map[string]interface{} {
	app := make(map[string]interface{})
	app["answers"] = flattenAnswers(ctx, &item.Answers)
	app["applied_at"] = item.AppliedAt
	app["attachments"] = flattenAttachments(ctx, &item.Attachments)
	app["candidate_id"] = item.CandidateId
	app["credited_to"] = flattenUser(ctx, item.CreditedTo)
	convertedStage := greenhouse.TypeIdName(*item.CurrentStage)
	app["current_stage"] = flattenTypeIdName(ctx, &convertedStage)
	app["custom_fields"] = item.CustomFields
	app["jobs"] = flattenJobs(ctx, &item.Jobs)
	app["job_post_id"] = item.JobPostId
	app["keyed_custom_fields"] = item.KeyedCustomFields
	app["last_activity_at"] = item.LastActivityAt
	app["location"] = flattenLocation(ctx, item.Location)
	app["prospect"] = item.Prospect
	app["prospect_detail"] = flattenProspectDetail(ctx, item.ProspectDetail)
	app["prospective_department"] = flattenDepartment(ctx, item.ProspectiveDepartment)
	app["prospective_office"] = flattenOffice(ctx, item.ProspectiveOffice)
	app["rejected_at"] = item.RejectedAt
	app["rejection_details"] = item.RejectionDetails
	app["rejection_reason"] = item.RejectionReason
	app["source"] = flattenSource(ctx, item.Source)
	app["status"] = item.Status
	return app
}

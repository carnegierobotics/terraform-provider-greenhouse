package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApplication() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"advance": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"answers": {
			Type:     schema.TypeList,
			Computed: true,
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
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
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
			Type:        schema.TypeList,
			Description: "The user who will receive credit for this application.",
			MaxItems:    1,
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
		"current_stage": {
			Type:     schema.TypeList,
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
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"hire": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"initial_stage_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"job_id": {
			Type: schema.TypeInt,
			//This is actually Required for candidates
			Optional: true,
		},
		"job_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"jobs": {
			Type:     schema.TypeList,
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
			Type:     schema.TypeList,
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
		"prospect": {
			Type:        schema.TypeBool,
			Description: "The candidate is a prospect and has not yet applied.",
			//This is actually required for Prospects
			Optional: true,
			Computed: true,
		},
		"prospect_detail": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseProspectDetail(),
			},
		},
		"prospect_owner_id": {
			Type:     schema.TypeInt,
			Optional: true,
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
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"prospective_department_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"prospective_office": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffice(),
			},
		},
		"prospective_office_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"referrer": {
			Type:     schema.TypeList,
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
		"rejection_details": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRejectionDetails(),
			},
		},
		"rejection_reason": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRejectionReason(),
			},
		},
		"source": {
			Type:        schema.TypeList,
			Description: "",
			Optional:    true,
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

func inflateApplications(ctx context.Context, list []interface{}) *[]greenhouse.Application {
	var newList []greenhouse.Application
	for i, item := range list {
		itemMap := item.(map[string]interface{})
		newList[i] = *inflateApplication(ctx, itemMap)
	}
	return &newList
}

func inflateApplication(ctx context.Context, item map[string]interface{}) *greenhouse.Application {
	var app greenhouse.Application
	app.Answers = *inflateAnswers(ctx, item["answers"])
	app.AppliedAt = item["applied_at"].(string)
	app.Attachments = *inflateAttachments(ctx, item["attachments"])
	app.CandidateId = item["candidate_id"].(int)
	inflatedCreditedTo := *inflateUsers(ctx, item["credited_to"])
	app.CreditedTo = &inflatedCreditedTo[0]
	inflatedCurrentStage := greenhouse.Stage(*inflateTypeIdName(ctx, item["current_stage"]))
	app.CurrentStage = &inflatedCurrentStage
	app.CustomFields = item["custom_fields"].(map[string]string)
	app.Id = item["id"].(int)
	app.InitialStageId = item["initial_stage_id"].(int)
	app.JobId = item["job_id"].(int)
	app.JobIds = item["job_ids"].([]int)
	app.Jobs = *inflateJobs(ctx, item["jobs"])
	app.JobPostId = item["job_post_id"].(int)
	app.KeyedCustomFields = *inflateKeyedCustomFields(ctx, item["keyed_custom_fields"])
	app.LastActivityAt = item["last_activity_at"].(string)
	app.Location = inflateLocation(ctx, item["location"])
	app.Prospect = item["prospect"].(bool)
	app.ProspectDetail = inflateProspectDetail(ctx, item["prospect_detail"])
	app.ProspectOwnerId = item["prospect_owner_id"].(int)
	app.ProspectPoolId = item["prospect_pool_id"].(int)
	app.ProspectPoolStageId = item["prospect_pool_stage_id"].(int)
	app.ProspectStageId = item["prospect_stage_id"].(int)
	app.ProspectiveDepartment = inflateDepartment(ctx, item["prospective_department"])
	app.ProspectiveDepartmentId = item["prospective_department_id"].(int)
	app.ProspectiveOffice = inflateOffice(ctx, item["prospective_office"])
	app.ProspectiveOfficeId = item["prospective_office_id"].(int)
	app.Referrer = inflateTypeTypeValue(ctx, item["referrer"])
	app.RejectedAt = item["rejected_at"].(string)
	app.RejectionDetails = inflateRejectionDetails(ctx, item["rejection_details"])
	app.RejectionReason = inflateRejectionReason(ctx, item["rejection_reason"])
	app.Source = inflateSource(ctx, item["source"])
	app.SourceId = item["source_id"].(int)
	app.Status = item["status"].(string)
	return &app
}

func flattenApplications(ctx context.Context, list *[]greenhouse.Application) []interface{} {
	if list != nil {
		tflog.Debug(ctx, "Flattening applications.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApplication(ctx, &item)
		}
		tflog.Debug(ctx, "Finished flattening applications.")
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApplication(ctx context.Context, item *greenhouse.Application) map[string]interface{} {
	tflog.Debug(ctx, "Flattening one application.")
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
	tflog.Debug(ctx, "Finished flattening application.")
	return app
}

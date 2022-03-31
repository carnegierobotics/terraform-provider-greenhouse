package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateApplications(ctx context.Context, source *[]interface{}) (*[]greenhouse.Application, diag.Diagnostics) {
	list := make([]greenhouse.Application, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		app, err := inflateApplication(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *app
	}
	return &list, nil
}

func inflateApplication(ctx context.Context, item *map[string]interface{}) (*greenhouse.Application, diag.Diagnostics) {
	var app greenhouse.Application
	tflog.Debug(ctx, "Inflating application.")
	if v, ok := (*item)["answers"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateAnswers(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Answers = *list
	}
	if v, ok := (*item)["applied_at"].(string); ok && len(v) > 0 {
		app.AppliedAt = v
	}
	if v, ok := (*item)["attachments"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating attachments.")
		list, err := inflateAttachments(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Attachments = *list
	}
	if v, ok := (*item)["candidate_id"].(int); ok {
		app.CandidateId = v
	}
	if v, ok := (*item)["credited_to"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating credited to.")
		inflatedCreditedTo, err := inflateUser(ctx, &v[0])
		if err != nil {
			return nil, err
		}
		if inflatedCreditedTo != nil {
			app.CreditedTo = inflatedCreditedTo
		}
	}
	if v, ok := (*item)["current_stage"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating current stage.")
		inflatedStage, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		convertedStage := greenhouse.Stage((*inflatedStage)[0])
		app.CurrentStage = &convertedStage
	}
	if v, ok := (*item)["custom_fields"].(map[string]interface{}); ok && len(v) > 0 {
		app.CustomFields = *mapAItoMapAA(ctx, v)
	}
	if v, ok := (*item)["id"].(int); ok {
		app.Id = v
	}
	if v, ok := (*item)["initial_stage_id"].(int); ok {
		app.InitialStageId = v
	}
	if v, ok := (*item)["job_id"].(int); ok {
		app.JobId = v
	}
	if v, ok := (*item)["job_ids"].([]interface{}); ok && len(v) > 0 {
		jobIds := v
		app.JobIds = *sliceItoSliceD(&jobIds)
	}
	if v, ok := (*item)["jobs"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating jobs.")
		list, err := inflateJobs(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Jobs = *list
	}
	if v, ok := (*item)["job_post_id"].(int); ok {
		app.JobPostId = v
	}
	if v, ok := (*item)["keyed_custom_fields"].(map[string]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating keyed custom fields.")
		app.KeyedCustomFields = *inflateKeyedCustomFields(ctx, &v)
	}
	if v, ok := (*item)["last_activity_at"].(string); ok && len(v) > 0 {
		app.LastActivityAt = v
	}
	if v, ok := (*item)["location"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating location.")
		list, err := inflateLocations(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Location = &(*list)[0]
	}
	if v, ok := (*item)["prospect"].(bool); ok {
		app.Prospect = v
	}
	if v, ok := (*item)["prospect_detail"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating prospect detail.")
		list, err := inflateProspectDetails(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.ProspectDetail = &(*list)[0]
	}
	if v, ok := (*item)["prospect_owner_id"].(int); ok {
		app.ProspectOwnerId = v
	}
	if v, ok := (*item)["prospect_pool_id"].(int); ok {
		app.ProspectPoolId = v
	}
	if v, ok := (*item)["prospect_pool_stage_id"].(int); ok {
		app.ProspectPoolStageId = v
	}
	if v, ok := (*item)["prospect_stage_id"].(int); ok {
		app.ProspectStageId = v
	}
	if v, ok := (*item)["prospective_department"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating department.")
		list, err := inflateDepartments(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.ProspectiveDepartment = &(*list)[0]
	}
	if v, ok := (*item)["prospective_department_id"].(int); ok {
		app.ProspectiveDepartmentId = v
	}
	if v, ok := (*item)["prospective_office"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating office.")
		list, err := inflateOffices(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.ProspectiveOffice = &(*list)[0]
	}
	if v, ok := (*item)["prospective_office_id"].(int); ok {
		app.ProspectiveOfficeId = v
	}
	if v, ok := (*item)["referrer"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating referrer.")
		inflatedReferrer, err := inflateTypeTypeValues(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Referrer = &(*inflatedReferrer)[0]
	}
	if v, ok := (*item)["rejected_at"].(string); ok && len(v) > 0 {
		app.RejectedAt = v
	}
	if v, ok := (*item)["rejection_details"].([]interface{}); ok && len(v) > 0 {
		tflog.Debug(ctx, "Inflating rejection details.")
		app.RejectionDetails = inflateRejectionDetails(ctx, v[0])
	}
	if v, ok := (*item)["rejection_reason"].([]interface{}); ok && len(v) > 0 {
		app.RejectionReason = inflateRejectionReason(ctx, v[0])
	}
	if v, ok := (*item)["source"].([]interface{}); ok && len(v) > 0 {
		app.Source = inflateSource(ctx, v)
	}
	if v, ok := (*item)["source_id"].(int); ok {
		app.SourceId = v
	}
	if v, ok := (*item)["status"].(string); ok && len(v) > 0 {
		app.Status = v
	}
	tflog.Debug(ctx, "Done inflating application.")
	return &app, nil
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

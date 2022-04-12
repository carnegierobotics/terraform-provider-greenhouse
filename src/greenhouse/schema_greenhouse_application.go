/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of terraform-provider-greenhouse.

terraform-provider-greenhouse is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

terraform-provider-greenhouse is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with terraform-provider-greenhouse. If not, see <https://www.gnu.org/licenses/>.
*/
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
			Optional:    true,
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
	tflog.Trace(ctx, "Inflating application.")
	if v, ok := (*item)["answers"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateAnswers(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Answers = *list
	}
	if v, ok := (*item)["applied_at"].(string); ok && len(v) > 0 {
		app.AppliedAt = &v
	}
	if v, ok := (*item)["attachments"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating attachments.")
		list, err := inflateAttachments(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Attachments = *list
	}
	if v, ok := (*item)["candidate_id"].(int); ok {
		app.CandidateId = &v
	}
	if v, ok := (*item)["credited_to"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating credited to.")
		list, err := inflateUsers(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.CreditedTo = &(*list)[0]
	}
	if v, ok := (*item)["current_stage"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating current stage.")
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
		app.Id = &v
	}
	if v, ok := (*item)["initial_stage_id"].(int); ok {
		app.InitialStageId = &v
	}
	if v, ok := (*item)["job_id"].(int); ok {
		app.JobId = &v
	}
	if v, ok := (*item)["job_ids"].([]interface{}); ok && len(v) > 0 {
		jobIds := &v
		app.JobIds = *sliceItoSliceD(jobIds)
	}
	if v, ok := (*item)["jobs"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating jobs.")
		list, err := inflateJobs(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Jobs = *list
	}
	if v, ok := (*item)["job_post_id"].(int); ok {
		app.JobPostId = &v
	}
	if v, ok := (*item)["keyed_custom_fields"].(map[string]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating keyed custom fields.")
		fields, err := inflateKeyedCustomFields(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.KeyedCustomFields = *fields
	}
	if v, ok := (*item)["last_activity_at"].(string); ok && len(v) > 0 {
		app.LastActivityAt = &v
	}
	if v, ok := (*item)["location"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating location.")
		list, err := inflateLocations(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Location = &(*list)[0]
	}
	if v, ok := (*item)["prospect"].(bool); ok {
		app.Prospect = &v
	}
	if v, ok := (*item)["prospect_detail"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating prospect detail.")
		list, err := inflateProspectDetails(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.ProspectDetail = &(*list)[0]
	}
	if v, ok := (*item)["prospect_owner_id"].(int); ok {
		app.ProspectOwnerId = &v
	}
	if v, ok := (*item)["prospect_pool_id"].(int); ok {
		app.ProspectPoolId = &v
	}
	if v, ok := (*item)["prospect_pool_stage_id"].(int); ok {
		app.ProspectPoolStageId = &v
	}
	if v, ok := (*item)["prospect_stage_id"].(int); ok {
		app.ProspectStageId = &v
	}
	if v, ok := (*item)["prospective_department"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating department.")
		list, err := inflateDepartments(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.ProspectiveDepartment = &(*list)[0]
	}
	if v, ok := (*item)["prospective_department_id"].(int); ok {
		app.ProspectiveDepartmentId = &v
	}
	if v, ok := (*item)["prospective_office"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating office.")
		list, err := inflateOffices(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.ProspectiveOffice = &(*list)[0]
	}
	if v, ok := (*item)["prospective_office_id"].(int); ok {
		app.ProspectiveOfficeId = &v
	}
	if v, ok := (*item)["referrer"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating referrer.")
		inflatedReferrer, err := inflateTypeTypeValues(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.Referrer = &(*inflatedReferrer)[0]
	}
	if v, ok := (*item)["rejected_at"].(string); ok && len(v) > 0 {
		app.RejectedAt = &v
	}
	if v, ok := (*item)["rejection_details"].([]interface{}); ok && len(v) > 0 {
		tflog.Trace(ctx, "Inflating rejection details.")
		list, err := inflateRejectionDetailsList(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.RejectionDetails = &(*list)[0]
	}
	if v, ok := (*item)["rejection_reason"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateRejectionReasons(ctx, &v)
		if err != nil {
			return nil, err
		}
		app.RejectionReason = &(*list)[0]
	}
	if v, ok := (*item)["source"].([]interface{}); ok && len(v) > 0 {
		source, ok := v[0].(map[string]interface{})
		if ok {
			obj, err := inflateSource(ctx, &source)
			if err != nil {
				return nil, err
			}
			app.Source = obj
		}
	}
	if v, ok := (*item)["source_id"].(int); ok {
		app.SourceId = &v
	}
	if v, ok := (*item)["status"].(string); ok && len(v) > 0 {
		app.Status = &v
	}
	tflog.Trace(ctx, "Done inflating application.")
	return &app, nil
}

func flattenApplications(ctx context.Context, list *[]greenhouse.Application) []interface{} {
	if list != nil {
		tflog.Trace(ctx, "Flattening applications.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApplication(ctx, &item)
		}
		tflog.Trace(ctx, "Finished flattening applications.")
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApplication(ctx context.Context, item *greenhouse.Application) map[string]interface{} {
	tflog.Trace(ctx, "Flattening one application.")
	app := make(map[string]interface{})
	app["answers"] = flattenAnswers(ctx, &item.Answers)
	if v := item.AppliedAt; v != nil {
		app["applied_at"] = *v
	}
	app["attachments"] = flattenAttachments(ctx, &item.Attachments)
	if v := item.CandidateId; v != nil {
		app["candidate_id"] = *v
	}
	if v := item.CreditedTo; v != nil {
		app["credited_to"] = flattenUser(ctx, v)
	} else {
		app["credited_to"] = emptyList()
	}
	if v := item.CurrentStage; v != nil {
		convertedStage := greenhouse.TypeIdName(*v)
		app["current_stage"] = []map[string]interface{}{flattenTypeIdName(ctx, &convertedStage)}
	} else {
		app["credited_to"] = emptyList()
	}
	app["custom_fields"] = item.CustomFields
	app["jobs"] = flattenJobs(ctx, &item.Jobs)
	if v := item.JobPostId; v != nil {
		app["job_post_id"] = *v
	}
	//app["keyed_custom_fields"] = item.KeyedCustomFields
	if v := item.LastActivityAt; v != nil {
		app["last_activity_at"] = *v
	}
	if v := item.Location; v != nil {
		app["location"] = flattenLocation(ctx, v)
	} else {
		app["location"] = emptyList()
	}
	if v := item.Prospect; v != nil {
		app["prospect"] = *v
	}
	if v := item.ProspectDetail; v != nil {
		app["prospect_detail"] = []map[string]interface{}{flattenProspectDetail(ctx, v)}
	} else {
		app["prospect_detail"] = emptyList()
	}
	if v := item.ProspectiveDepartment; v != nil {
		app["prospective_department"] = flattenDepartment(ctx, v)
	} else {
		app["prospective_department"] = emptyList()
	}
	if v := item.ProspectiveOffice; v != nil {
		app["prospective_office"] = flattenOffice(ctx, v)
	} else {
		app["prospective_office"] = emptyList()
	}
	if v := item.RejectedAt; v != nil {
		app["rejected_at"] = *v
	}
	if v := item.RejectionDetails; v != nil {
		app["rejection_details"] = flattenRejectionDetails(ctx, v)
	} else {
		app["rejection_details"] = emptyList()
	}
	if v := item.RejectionReason; v != nil {
		app["rejection_reason"] = flattenRejectionReason(ctx, v)
	}
	if v := item.Source; v != nil {
		app["source"] = []map[string]interface{}{flattenSource(ctx, v)}
	}
	if v := item.Status; v != nil {
		app["status"] = *v
	}
	tflog.Trace(ctx, "Finished flattening application.")
	return app
}

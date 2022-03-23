package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseApplicationCreate,
		ReadContext:   resourceGreenhouseApplicationRead,
		UpdateContext: resourceGreenhouseApplicationUpdate,
		DeleteContext: resourceGreenhouseApplicationDelete,
		Exists:        resourceGreenhouseApplicationExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseTypeIdName(),
	}
}

func resourceGreenhouseApplicationExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/applications/%d", id))
}

func resourceGreenhouseApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Create is not supported for applications."}}
}

func resourceGreenhouseApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetApplication(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.Set("answers", flattenAnswers(ctx, &obj.Answers))
	d.Set("applied_at", obj.AppliedAt)
	d.Set("attachments", flattenAttachments(ctx, &obj.Attachments))
	d.Set("candidate_id", obj.CandidateId)
	d.Set("credited_to", flattenUserBasics(ctx, obj.CreditedTo))
	convertedStage := greenhouse.TypeIdName(*obj.CurrentStage)
	d.Set("current_stage", flattenTypeIdName(ctx, &convertedStage))
	d.Set("custom_fields", obj.CustomFields)
	d.Set("job_post_id", obj.JobPostId)
	d.Set("jobs", flattenJobs(ctx, &obj.Jobs))
	d.Set("keyed_custom_fields", flattenKeyedCustomFields(ctx, &obj.KeyedCustomFields))
	d.Set("last_activity_at", obj.LastActivityAt)
	d.Set("location", flattenLocation(ctx, obj.Location))
	d.Set("prospect", obj.Prospect)
	d.Set("prospect_detail", flattenProspectDetail(ctx, obj.ProspectDetail))
	d.Set("prospective_department", flattenDepartment(ctx, obj.ProspectiveDepartment))
	d.Set("prospective_office", flattenOffice(ctx, obj.ProspectiveOffice))
	d.Set("rejected_at", obj.RejectedAt)
	d.Set("rejection_reason", obj.RejectionReason)
	d.Set("source", flattenSource(ctx, obj.Source))
	d.Set("status", obj.Status)
	return nil
}

func resourceGreenhouseApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for applications."}}
}

func resourceGreenhouseApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for applications."}}
}

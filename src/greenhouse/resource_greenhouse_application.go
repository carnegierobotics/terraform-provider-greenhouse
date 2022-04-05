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
			StateContext: resourceGreenhouseApplicationImport,
		},
		Schema: schemaGreenhouseApplication(),
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
	var createObj greenhouse.Application
	referrer := d.Get("referrer").([]interface{})
	if len(referrer) == 1 {
		referrerObj, err := inflateTypeTypeValues(ctx, &referrer)
		if err != nil {
			return err
		}
		if referrerObj != nil && len(*referrerObj) > 0 {
			createObj.Referrer = &(*referrerObj)[0]
		}
	}
	if d.Get("prospect").(bool) {
		createObj.Prospect = BoolPtr(d.Get("prospect").(bool))
		createObj.JobIds = d.Get("job_ids").([]int)
		createObj.SourceId = IntPtr(d.Get("source_id").(int))
		createObj.ProspectPoolId = IntPtr(d.Get("prospect_pool_id").(int))
		createObj.ProspectPoolStageId = IntPtr(d.Get("prospect_pool_stage_id").(int))
		createObj.ProspectOwnerId = IntPtr(d.Get("prospect_owner_id").(int))
		createObj.ProspectiveDepartmentId = IntPtr(d.Get("prospective_department_id").(int))
		createObj.ProspectiveOfficeId = IntPtr(d.Get("prospective_office_id").(int))
	} else {
		createObj.JobId = IntPtr(d.Get("job_id").(int))
		createObj.SourceId = IntPtr(d.Get("source_id").(int))
		createObj.InitialStageId = IntPtr(d.Get("initial_stage_id").(int))
		attachments := d.Get("attachments").([]interface{})
		obj, err := inflateAttachments(ctx, &attachments)
		if err != nil {
			return err
		}
		createObj.Attachments = *obj
	}
	return resourceGreenhouseApplicationUpdate(ctx, d, meta)
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
	for k, v := range flattenApplication(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var err error
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	updateObj := greenhouse.Application{
		SourceId:        IntPtr(d.Get("source_id").(int)),
		Referrer:        d.Get("referrer").(*greenhouse.TypeTypeValue),
		CustomFields:    d.Get("custom_fields").(map[string]string),
		ProspectPoolId:  IntPtr(d.Get("prospect_pool_id").(int)),
		ProspectStageId: IntPtr(d.Get("prospect_stage_id").(int)),
	}
	err = greenhouse.UpdateApplication(meta.(*greenhouse.Client), ctx, id, &updateObj)
	hire := d.Get("hire").(bool)
	reject := d.Get("reject").(bool)
	if hire == true && reject == true {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Hire and reject cannot both be true."}}
	}
	if !d.HasChanges("job_id") && d.HasChanges("current_stage") {
		from, to := d.GetChange("current_stage.0.id")
		err = greenhouse.MoveApplicationSameJob(meta.(*greenhouse.Client), ctx, id, from.(int), to.(int))
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	if d.HasChanges("hire") {
		if hire == true {
			hireObj := greenhouse.ApplicationHire{
				CloseReasonId: IntPtr(d.Get("close_reason_id").(int)),
				OpeningId:     IntPtr(d.Get("opening_id").(int)),
				StartDate:     StringPtr(d.Get("start_date").(string)),
			}
			err = greenhouse.HireApplication(meta.(*greenhouse.Client), ctx, id, &hireObj)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		} else {
			return diag.Diagnostics{{Severity: diag.Error, Summary: "Not possible to un-hire someone."}}
		}
	}
	if d.HasChanges("reject") {
		if reject == true {
			rejectionEmail := greenhouse.RejectionEmail{
				EmailTemplateId: StringPtr(d.Get("email_template_id").(string)),
				SendEmailAt:     StringPtr(d.Get("send_email_at").(string)),
			}
			rejectObj := greenhouse.ApplicationReject{
				Notes:             StringPtr(d.Get("notes").(string)),
				RejectionEmail:    &rejectionEmail,
				RejectionReasonId: IntPtr(d.Get("rejection_reason_id").(int)),
			}
			err = greenhouse.RejectApplication(meta.(*greenhouse.Client), ctx, id, &rejectObj)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		} else {
			err = greenhouse.UnrejectApplication(meta.(*greenhouse.Client), ctx, id)
		}
	}
	if d.HasChanges("rejection_reason") {
		rejId := Int(d.Get("rejection_reason.0").(greenhouse.RejectionReason).Id)
		err = greenhouse.UpdateRejectionReason(meta.(*greenhouse.Client), ctx, id, rejId)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	if d.Get("advance").(bool) == true {
		from := Int(d.Get("current_stage.0").(greenhouse.TypeIdName).Id)
		err = greenhouse.AdvanceApplication(meta.(*greenhouse.Client), ctx, id, from)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	if d.HasChanges("job_id") {
		newJob := d.Get("job_id").(int)
		newStage := Int(d.Get("current_stage.0").(greenhouse.TypeIdName).Id)
		err = greenhouse.MoveApplicationDifferentJob(meta.(*greenhouse.Client), ctx, id, newJob, newStage)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return resourceGreenhouseApplicationRead(ctx, d, meta)
}

func resourceGreenhouseApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteApplication(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("")
	return nil
}

func resourceGreenhouseApplicationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseApplicationRead)
}

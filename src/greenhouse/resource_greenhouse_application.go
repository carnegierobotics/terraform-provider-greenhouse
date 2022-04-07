package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	//"github.com/hashicorp/terraform-plugin-log/tflog"
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
	var obj greenhouse.Application
	referrer := d.Get("referrer").([]interface{})
	if len(referrer) == 1 {
		referrerObj, err := inflateTypeTypeValues(ctx, &referrer)
		if err != nil {
			return err
		}
		if referrerObj != nil && len(*referrerObj) > 0 {
			obj.Referrer = &(*referrerObj)[0]
		}
	}
	if v, ok := d.Get("source_id").(int); ok {
		obj.SourceId = &v
	}
	if v, ok := d.Get("prospect").(bool); ok && v {
		obj.Prospect = &v
		if v, ok := d.Get("job_ids").([]int); ok && len(v) > 0 {
			obj.JobIds = v
		}
		if v, ok := d.Get("prospect_pool_id").(int); ok {
			obj.ProspectPoolId = &v
		}
		if v, ok := d.Get("prospect_pool_stage_id").(int); ok {
			obj.ProspectPoolStageId = &v
		}
		if v, ok := d.Get("prospect_owner_id").(int); ok {
			obj.ProspectOwnerId = &v
		}
		if v, ok := d.Get("prospective_department_id").(int); ok {
			obj.ProspectiveDepartmentId = &v
		}
		if v, ok := d.Get("prospective_office_id").(int); ok {
			obj.ProspectiveOfficeId = &v
		}
	} else {
		if v, ok := d.Get("job_id").(int); ok {
			obj.JobId = &v
		}
		if v, ok := d.Get("initial_stage_id").(int); ok {
			obj.InitialStageId = &v
		}
		if v, ok := d.Get("attachments").([]interface{}); ok && len(v) > 0 {
			attachObj, err := inflateAttachments(ctx, &v)
			if err != nil {
				return err
			}
			obj.Attachments = *attachObj
		}
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
	var obj greenhouse.Application
	if v, ok := d.Get("source_id").(int); ok {
		obj.SourceId = &v
	}
	if v, ok := d.Get("referrer").(*greenhouse.TypeTypeValue); ok {
		obj.Referrer = v
	}
	if v, ok := d.Get("custom_fields").(map[string]string); ok && len(v) > 0 {
		obj.CustomFields = v
	}
	if v, ok := d.Get("prospect_pool_id").(int); ok && v != 0 {
		obj.ProspectPoolId = &v
	}
	if v, ok := d.Get("prospect_stage_id").(int); ok && v != 0 {
		obj.ProspectStageId = &v
	}
	diagErr := logJson(ctx, "resourceGreenhouseApplicationUpdate", obj)
	if diagErr != nil {
		return diagErr
	}
	err = greenhouse.UpdateApplication(meta.(*greenhouse.Client), ctx, id, &obj)
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
			var hireObj greenhouse.ApplicationHire
			if v, ok := d.Get("close_reason_id").(int); ok {
				hireObj.CloseReasonId = &v
			}
			if v, ok := d.Get("opening_id").(int); ok {
				hireObj.OpeningId = &v
			}
			if v, ok := d.Get("start_date").(string); ok && len(v) > 0 {
				hireObj.StartDate = &v
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
		if reject {
			var rejectionEmail greenhouse.RejectionEmail
			if v, ok := d.Get("email_template_id").(string); ok && len(v) > 0 {
				rejectionEmail.EmailTemplateId = &v
			}
			if v, ok := d.Get("send_email_at").(string); ok && len(v) > 0 {
				rejectionEmail.SendEmailAt = &v
			}
			var rejectObj greenhouse.ApplicationReject
			if v, ok := d.Get("notes").(string); ok && len(v) > 0 {
				rejectObj.Notes = &v
			}
			rejectObj.RejectionEmail = &rejectionEmail
			if v, ok := d.Get("rejection_reason_id").(int); ok {
				rejectObj.RejectionReasonId = &v
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
		if v, ok := d.Get("rejection_reason.0").(greenhouse.RejectionReason); ok {
			rejId := *v.Id
			err = greenhouse.UpdateRejectionReason(meta.(*greenhouse.Client), ctx, id, rejId)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		}
	}
	if v, ok := d.Get("advance").(bool); ok && v {
		if v, ok := d.Get("current_stage.0").(greenhouse.TypeIdName); ok && v.Id != nil {
			from := *v.Id
			err = greenhouse.AdvanceApplication(meta.(*greenhouse.Client), ctx, id, from)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
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

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
	hire := d.Get("hire").(bool)
	reject := d.Get("reject").(bool)
	if hire == true && reject == true {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Hire and reject cannot both be true."}}
	}
	if d.HasChanges("hire") {
		if hire == true {
			hireObj := greenhouse.ApplicationHire{CloseReasonId: d.Get("close_reason_id").(int), OpeningId: d.Get("opening_id").(int), StartDate: d.Get("start_date").(string)}
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
			rejectionEmail := greenhouse.RejectionEmail{EmailTemplateId: d.Get("email_template_id").(string), SendEmailAt: d.Get("send_email_at").(string)}
			rejectObj := greenhouse.ApplicationReject{Notes: d.Get("notes").(string), RejectionEmail: &rejectionEmail, RejectionReasonId: d.Get("rejection_reason_id").(int)}
			err = greenhouse.RejectApplication(meta.(*greenhouse.Client), ctx, id, &rejectObj)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		} else {
			err = greenhouse.UnrejectApplication(meta.(*greenhouse.Client), ctx, id)
		}
	}
	return nil
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

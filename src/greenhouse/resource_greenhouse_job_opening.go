package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseJobOpening() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseJobOpeningCreate,
		ReadContext:   resourceGreenhouseJobOpeningRead,
		UpdateContext: resourceGreenhouseJobOpeningUpdate,
		DeleteContext: resourceGreenhouseJobOpeningDelete,
		Exists:        resourceGreenhouseJobOpeningExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseJobOpening(),
	}
}

func resourceGreenhouseJobOpeningExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/job_openings/%d", id))
}

func resourceGreenhouseJobOpeningCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Create is not supported for job_openings."}}
}

func resourceGreenhouseJobOpeningRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if v, ok := d.Get("job_id").(int); ok {
		obj, err := greenhouse.GetJobOpening(meta.(*greenhouse.Client), ctx, v, id)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		d.Set("application_id", obj.ApplicationId)
		d.Set("close_reason", flattenCloseReason(ctx, obj.CloseReason))
		return nil
	}
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting job_opening."}}
}

func resourceGreenhouseJobOpeningUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for job_openings."}}
}

func resourceGreenhouseJobOpeningDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for job_openings."}}
}

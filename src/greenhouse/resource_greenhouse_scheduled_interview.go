package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseScheduledInterview() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseScheduledInterviewCreate,
		ReadContext:   resourceGreenhouseScheduledInterviewRead,
		UpdateContext: resourceGreenhouseScheduledInterviewUpdate,
		DeleteContext: resourceGreenhouseScheduledInterviewDelete,
		Exists:        resourceGreenhouseScheduledInterviewExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseScheduledInterview(),
	}
}

func resourceGreenhouseScheduledInterviewExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/scheduled_interviews/%d", id))
}

func resourceGreenhouseScheduledInterviewCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Create is not supported for xxx."}}
}

func resourceGreenhouseScheduledInterviewRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetScheduledInterview(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenScheduledInterview(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseScheduledInterviewUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	updateObj := greenhouse.ScheduledInterviewUpdateInfo{End: d.Get("end").(string), ExternalEventId: d.Get("external_event_id").(string)}
	err = greenhouse.UpdateScheduledInterview(meta.(*greenhouse.Client), ctx, id, &updateObj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return nil
}

func resourceGreenhouseScheduledInterviewDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteScheduledInterview(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("")
	return nil
}

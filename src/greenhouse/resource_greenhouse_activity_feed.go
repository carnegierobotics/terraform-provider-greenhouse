package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseActivityFeed() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseActivityFeedCreate,
		ReadContext:   resourceGreenhouseActivityFeedRead,
		UpdateContext: resourceGreenhouseActivityFeedUpdate,
		DeleteContext: resourceGreenhouseActivityFeedDelete,
		Exists:        resourceGreenhouseActivityFeedExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseTypeIdName(),
	}
}

func resourceGreenhouseActivityFeedExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/activity_feed/%d", id))
}

func resourceGreenhouseActivityFeedCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Create is not supported for activity_feed."}}
}

func resourceGreenhouseActivityFeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetActivityFeed(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.Set("activities", flattenActivities(ctx, &obj.Activities))
	d.Set("emails", flattenEmails(ctx, &obj.Emails))
	d.Set("notes", flattenNotes(ctx, &obj.Notes))
	return nil
}

func resourceGreenhouseActivityFeedUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for activity_feed."}}
}

func resourceGreenhouseActivityFeedDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for activity_feed."}}
}

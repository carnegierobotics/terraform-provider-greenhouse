package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseActivityFeed() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseActivityFeedRead,
		Schema: schemaGreenhouseActivityFeed(),
	}
}

func dataSourceGreenhouseActivityFeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	candidateId := d.Get("candidate_id").(int)
	obj, err := greenhouse.GetActivityFeed(meta.(*greenhouse.Client), ctx, candidateId)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  d.SetId(strconv.Itoa(candidateId))
	d.Set("activities", flattenActivities(ctx, &obj.Activities))
	d.Set("emails", flattenEmails(ctx, &obj.Emails))
	d.Set("notes", flattenNotes(ctx, &obj.Notes))
	return nil
}

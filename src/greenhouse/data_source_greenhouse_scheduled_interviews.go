package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseScheduledInterviews() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseScheduledInterviewsRead,
		Schema: map[string]*schema.Schema{
      "actionable": {
        Type: schema.TypeBool,
        Optional: true,
      },
      "application_id": {
        Type: schema.TypeInt,
        Optional: true,
      },
      "interviews": {
        Type: schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: schemaGreenhouseScheduledInterview(),
        },
      },
    },
	}
}

func dataSourceGreenhouseScheduledInterviewsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  appId, ok := d.GetOk("application_id")
  actionable := d.Get("actionable").(bool)
  var interviews *[]greenhouse.ScheduledInterview
  var err error
  var id string
  if ok {
    id = strconv.Itoa(appId.(int))
    interviews, err = greenhouse.GetScheduledInterviewsForApplication(meta.(*greenhouse.Client), ctx, appId.(int), actionable)
  } else {
    id = "all"
    interviews, err = greenhouse.GetAllScheduledInterviews(meta.(*greenhouse.Client), ctx, actionable)
  }
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  d.SetId(id)
	d.Set("interviews", flattenScheduledInterviews(ctx, interviews))
	return nil
}

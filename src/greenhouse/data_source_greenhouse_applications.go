package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseApplications() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseApplicationsRead,
		Schema: map[string]*schema.Schema{
      "applications": {
        Type: schema.TypeList,
        Computed: true,
        Elem: &schema.Resource{
          Schema: schemaGreenhouseApplication(),
        },
      },
      "created_after": {
        Type: schema.TypeString,
        Optional: true,
      },
      "created_before": {
        Type: schema.TypeString,
        Optional: true,
      },
      "job_id": {
        Type: schema.TypeInt,
        Optional: true,
      },
      "last_activity_after": {
        Type: schema.TypeString,
        Optional: true,
      },
      "status": {
        Type: schema.TypeString,
        Optional: true,
      },
    },
	}
}

func dataSourceGreenhouseApplicationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllApplications(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  d.SetId("all")
	d.Set("applications", flattenApplications(ctx, list))
	return nil
}

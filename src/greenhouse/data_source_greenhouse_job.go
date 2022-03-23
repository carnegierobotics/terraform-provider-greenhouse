package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseJob() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseJobRead,
		Schema: map[string]*schema.Schema{
      "name": {
        Type: schema.TypeString,
        Optional: true,
      },
    },
	}
}

func dataSourceGreenhouseJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllJobs(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  name := d.Get("name").(string)
  for _, job := range *list {
    if job.Name == name {
      d.SetId(strconf.Itoa(job.Id))
      return nil
    }
  }
	return nil
}

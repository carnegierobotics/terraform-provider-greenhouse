package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseJobs() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseJobsRead,
		Schema: map[string]*schema.Schema{
      "names": {
        Type: schema.TypeList,
        Computed: true,
        Elem: &schema.Schema{
          Type: schema.TypeString,
        },
      },
    },
	}
}

func dataSourceGreenhouseJobsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllJobs(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  jobs := make([]string, len(*list), len(*list))
  for i, job := range *list {
    jobs[i] = job.Name
  }
  d.SetId("all")
	d.Set("names", jobs)
	return nil
}

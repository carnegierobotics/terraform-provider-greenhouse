package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseJobStage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseJobStageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseJobStageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, ok := d.GetOk("id")
	var stage *greenhouse.JobStage
	var err error
	if ok {
		stage, err = greenhouse.GetJobStage(meta.(*greenhouse.Client), ctx, id.(int))
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	} else {
		job_id := d.Get("job_id").(int)
		name := d.Get("name").(string)
		list, err := greenhouse.GetJobStagesForJob(meta.(*greenhouse.Client), ctx, job_id)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		for _, item := range *list {
			if name == *item.Name {
				stage = &item
			}
		}
	}
	d.SetId(strconv.Itoa(*stage.Id))
	d.Set("created_at", stage.CreatedAt)
	d.Set("interviews", flattenInterviews(ctx, &stage.Interviews))
	d.Set("job_id", stage.JobId)
	d.Set("name", stage.Name)
	d.Set("priority", stage.Priority)
	d.Set("updated_at", stage.UpdatedAt)
	return nil
}

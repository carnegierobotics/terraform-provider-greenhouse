package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseJobOpenings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseJobOpeningsRead,
		Schema: map[string]*schema.Schema{
			"job_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"openings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseJobOpening(),
				},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGreenhouseJobOpeningsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	jobId := d.Get("job_id").(int)
	status := d.Get("status").(string)
	list, err := greenhouse.GetAllJobOpenings(meta.(*greenhouse.Client), ctx, jobId, status)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	id := strconv.Itoa(jobId)
	if status != "" {
		id = fmt.Sprintf("%s-%s", id, status)
	}
	d.SetId(id)
	d.Set("openings", flattenJobOpenings(ctx, list))
	return nil
}

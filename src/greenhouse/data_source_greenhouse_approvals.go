package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseApprovals() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseApprovalsRead,
		Schema: map[string]*schema.Schema{
			"approvals": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseApproval(),
				},
			},
			"job_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseApprovalsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("job_id").(int)
	list, err := greenhouse.ListApprovalsForJob(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	d.Set("approvals", flattenApprovals(ctx, list))
	return nil
}

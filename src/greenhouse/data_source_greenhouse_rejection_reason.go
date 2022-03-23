package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseRejectionReason() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseRejectionReasonRead,
		Schema:      schemaGreenhouseRejectionReason(),
	}
}

func dataSourceGreenhouseRejectionReasonRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	defaults := d.Get("include_defaults").(bool)
	pp := d.Get("per_page").(int)
	list, err := greenhouse.GetAllRejectionReasons(meta.(*greenhouse.Client), ctx, defaults, pp)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for _, reason := range *list {
		if reason.Name == name {
			d.SetId(strconv.Itoa(reason.Id))
			return nil
		}
	}
	return nil
}

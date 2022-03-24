package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseCloseReasons() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCloseReasonsRead,
		Schema:      schemaGreenhouseCloseReasons(),
	}
}

func dataSourceGreenhouseCloseReasonsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	reasons, err := greenhouse.GetAllCloseReasons(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("all")
	d.Set("reasons", flattenCloseReasons(ctx, reasons))
	return nil
}

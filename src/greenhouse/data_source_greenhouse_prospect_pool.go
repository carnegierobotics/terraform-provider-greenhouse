package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "strconv"
)

func dataSourceGreenhouseProspectPool() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseProspectPoolRead,
		Schema: schemaGreenhouseProspectPool(),
	}
}

func dataSourceGreenhouseProspectPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  name := d.Get("name").(string)
	list, err := greenhouse.GetAllProspectPools(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  for _, pool := range *list {
    if pool.Name == name {
      d.SetId(strconf.Itoa(pool.Id))
      return nil
    }
  }
	return nil
}

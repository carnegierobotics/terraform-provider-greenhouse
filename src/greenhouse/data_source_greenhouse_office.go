package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseOffice() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseOfficeRead,
		Schema: map[string]*schema.Schema{
      "name": {
        Type: schema.TypeString,
        Optional: true,
      },
    },
	}
}

func dataSourceGreenhouseOfficeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  name := d.Get("name").(string)
  if name != "" {
	  list, err := greenhouse.GetAllOffices(meta.(*greenhouse.Client), ctx)
	  if err != nil {
		  return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	  }
    for _, office := range *list {
      if office.Name == name {
        d.SetId(strconf.Itoa(office.Id))
        return nil
      }
    }
    return nil
  }
	return nil
}

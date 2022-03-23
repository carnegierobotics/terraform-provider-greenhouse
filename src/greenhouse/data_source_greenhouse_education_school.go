package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseSchool() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseSchoolRead,
		Schema: map[string]*schema.Schema{
      "name": {
        Type: schema.TypeString,
        Required: true,
      },
    },
	}
}

func dataSourceGreenhouseSchoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllSchools(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  name := d.Get("name").(string)
  for _, school := range *list {
    if school.Name == name {
      d.SetId(strconf.Itoa(school.Id))
      return nil
    }
  }
	return nil
}

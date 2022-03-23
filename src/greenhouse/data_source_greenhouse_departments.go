package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseDepartments() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseDepartmentsRead,
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

func dataSourceGreenhouseDepartmentsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllDepartments(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  departments := make([]string, len(*list), len(*list))
  for i, department := range *list {
    departments[i] = department.Name
  }
  d.SetId("all")
	d.Set("names", departments)
	return nil
}

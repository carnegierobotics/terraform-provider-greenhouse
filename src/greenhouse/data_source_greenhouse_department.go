package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseDepartment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDepartmentRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGreenhouseDepartmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllDepartments(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	name := d.Get("name").(string)
	for _, department := range *list {
		if *department.Name == name {
			d.SetId(strconv.Itoa(*department.Id))
			d.Set("child_department_external_ids", department.ChildDepartmentExternalIds)
			d.Set("child_ids", department.ChildIds)
			d.Set("parent_department_external_id", department.ParentDepartmentExternalId)
			d.Set("parent_id", department.ParentId)
			return nil
		}
	}
	return nil
}

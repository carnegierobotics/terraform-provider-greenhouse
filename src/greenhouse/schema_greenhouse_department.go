package greenhouse

import (
	"context"
  "fmt"
  "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDepartment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"child_department_external_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"child_ids": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"external_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"parent_department_external_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
}

func flattenDepartments(ctx context.Context, list *[]greenhouse.Department) []interface{} {
  tflog.Debug(ctx, "Flattening department list", "deptlist", fmt.Sprintf("%+v", list))
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenDepartment(ctx, &item)
		}
    tflog.Debug(ctx, "Flattened department list", "deptlist", fmt.Sprintf("%+v", flatList))
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenDepartment(ctx context.Context, item *greenhouse.Department) map[string]interface{} {
  tflog.Debug(ctx, "Flattening department", "department", fmt.Sprintf("%+v", item))
	dept := make(map[string]interface{})
	dept["name"] = item.Name
	dept["parent_id"] = item.ParentId
	dept["child_ids"] = item.ChildIds
  tflog.Debug(ctx, "Flattened department", "department", fmt.Sprintf("%+v", dept))
	return dept
}

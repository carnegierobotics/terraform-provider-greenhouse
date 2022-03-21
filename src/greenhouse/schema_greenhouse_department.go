package greenhouse

import (
	"context"
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
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenDepartment(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenDepartment(ctx context.Context, item *greenhouse.Department) map[string]interface{} {
	dept := make(map[string]interface{})
	dept["name"] = item.Name
	dept["parent_id"] = item.ParentId
	dept["child_ids"] = item.ChildIds
	return dept
}

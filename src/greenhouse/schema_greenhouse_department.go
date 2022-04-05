package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDepartment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"child_department_external_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"child_ids": {
			Type:     schema.TypeList,
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
		},
	}
}

func inflateDepartments(ctx context.Context, source *[]interface{}) (*[]greenhouse.Department, diag.Diagnostics) {
	list := make([]greenhouse.Department, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateDepartment(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateDepartment(ctx context.Context, source *map[string]interface{}) (*greenhouse.Department, diag.Diagnostics) {
	var obj greenhouse.Department
	if v, ok := (*source)["child_department_external_ids"].([]interface{}); ok && len(v) > 0 {
		obj.ChildDepartmentExternalIds = *sliceItoSliceA(&v)
	}
	if v, ok := (*source)["child_ids"].([]interface{}); ok && len(v) > 0 {
		obj.ChildIds = *sliceItoSliceD(&v)
	}
	if v, ok := (*source)["external_id"].(string); ok && len(v) > 0 {
		obj.ExternalId = &v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["parent_department_external_ids"].(string); ok && len(v) > 0 {
		obj.ParentDepartmentExternalId = &v
	}
	if v, ok := (*source)["parent_id"].(int); ok {
		obj.ParentId = &v
	}
	return &obj, nil
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
	if v := item.ChildDepartmentExternalIds; len(v) > 0 {
		dept["child_department_external_ids"] = v
	}
	if v := item.ChildIds; len(v) > 0 {
		dept["child_ids"] = v
	}
	if v := item.Name; v != nil {
		dept["name"] = *v
	}
	if v := item.ParentDepartmentExternalId; v != nil {
		dept["parent_department_external_id"] = *v
	}
	if v := item.ParentId; v != nil {
		dept["parent_id"] = *v
	}
	tflog.Debug(ctx, "Flattened department", "department", fmt.Sprintf("%+v", dept))
	return dept
}

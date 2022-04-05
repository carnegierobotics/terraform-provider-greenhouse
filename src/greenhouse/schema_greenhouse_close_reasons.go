package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCloseReasons() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"reasons": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func inflateCloseReasons(ctx context.Context, source *[]interface{}) (*[]greenhouse.CloseReason, diag.Diagnostics) {
	list := make([]greenhouse.CloseReason, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateCloseReason(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateCloseReason(ctx context.Context, source *map[string]interface{}) (*greenhouse.CloseReason, diag.Diagnostics) {
	var obj greenhouse.CloseReason
	if v, ok := (*source)["reasons"].([]interface{}); ok && len(v) > 0 {
		item, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		converted := greenhouse.CloseReason((*item)[0])
		obj = converted
	}
	return &obj, nil
}

func flattenCloseReasons(ctx context.Context, list *[]greenhouse.CloseReason) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			converted := greenhouse.TypeIdName(item)
			flatList[i] = flattenTypeIdName(ctx, &converted)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenCloseReason(ctx context.Context, item *greenhouse.CloseReason) map[string]interface{} {
	tflog.Debug(ctx, "Flattening close reason", "reason", fmt.Sprintf("%+v", item))
	flatItem := make(map[string]interface{})
	if v := item.Name; v != nil {
		flatItem["name"] = *v
	}
	tflog.Debug(ctx, "Flattened close reason", "reason", fmt.Sprintf("%+v", flatItem))
	return flatItem
}

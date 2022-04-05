package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseTypeTypeValue() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func inflateTypeTypeValues(ctx context.Context, source *[]interface{}) (*[]greenhouse.TypeTypeValue, diag.Diagnostics) {
	list := make([]greenhouse.TypeTypeValue, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateTypeTypeValue(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateTypeTypeValue(ctx context.Context, source *map[string]interface{}) (*greenhouse.TypeTypeValue, diag.Diagnostics) {
	var item greenhouse.TypeTypeValue
	if v, ok := (*source)["type"].(string); ok && len(v) > 0 {
		item.Type = &v
	}
	if v, ok := (*source)["value"].(string); ok && len(v) > 0 {
		item.Value = &v
	}
	return &item, nil
}

func flattenTypeTypeValues(ctx context.Context, list *[]greenhouse.TypeTypeValue) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenTypeTypeValue(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenTypeTypeValue(ctx context.Context, item *greenhouse.TypeTypeValue) map[string]interface{} {
	obj := make(map[string]interface{})
	if v := item.Type; v != nil {
		obj["type"] = *v
	}
	if v := item.Value; v != nil {
		obj["value"] = *v
	}
	return obj
}

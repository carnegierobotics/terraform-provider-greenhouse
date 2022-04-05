package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOpening() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_fields": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func inflateOpenings(ctx context.Context, source *[]interface{}) (*[]greenhouse.Opening, diag.Diagnostics) {
	list := make([]greenhouse.Opening, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateOpening(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateOpening(ctx context.Context, source *map[string]interface{}) (*greenhouse.Opening, diag.Diagnostics) {
	var obj greenhouse.Opening
	if v, ok := (*source)["custom_fields"].([]interface{}); ok && len(v) > 0 {
		list := make([]map[string]string, len(v), len(v))
		for i, item := range v {
			list[i] = make(map[string]string)
			for k, v := range item.(map[string]interface{}) {
				list[i][k] = v.(string)
			}
		}
		obj.CustomFields = list
	}
	return &obj, nil
}

func flattenOpenings(ctx context.Context, list *[]greenhouse.Opening) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenOpening(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenOpening(ctx context.Context, item *greenhouse.Opening) map[string]interface{} {
	opening := make(map[string]interface{})
	if v := item.CustomFields; len(v) > 0 {
		opening["custom_fields"] = v
	}
	return opening
}

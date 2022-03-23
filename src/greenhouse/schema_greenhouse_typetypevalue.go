package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseTypeTypeValue() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func convertToTypeTypeValueList(list []interface{}) *[]greenhouse.TypeTypeValue {
	newList := make([]greenhouse.TypeTypeValue, len(list))
	for i := range list {
		newList[i] = list[i].(greenhouse.TypeTypeValue)
	}
	return &newList
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
	obj["type"] = item.Type
	obj["value"] = item.Value
	return obj
}

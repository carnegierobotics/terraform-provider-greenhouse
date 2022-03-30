package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateTypeTypeValues(ctx context.Context, source *[]interface{}) (*[]greenhouse.TypeTypeValue, diag.Diagnostics) {
  if source != nil && len(*source) > 0 {
	  newList := make([]greenhouse.TypeTypeValue, len(*source))
	  for i, item := range *source {
		  newList[i] = item.(greenhouse.TypeTypeValue)
	  }
	  return &newList, nil
  }
  return nil, nil
}

func inflateTypeTypeValue(ctx context.Context, source interface{}) (*greenhouse.TypeTypeValue, diag.Diagnostics) {
	var item greenhouse.TypeTypeValue
	err := convertType(ctx, source, item)
  if err != nil {
    return nil, err
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
	obj["type"] = item.Type
	obj["value"] = item.Value
	return obj
}

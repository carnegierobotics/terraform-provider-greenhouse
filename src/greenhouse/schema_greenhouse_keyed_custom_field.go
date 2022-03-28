package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseKeyedCustomField() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
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

func inflateKeyedCustomFields(ctx context.Context, source interface{}) *map[string]greenhouse.KeyedCustomField {
	var list map[string]greenhouse.KeyedCustomField
	convertType(ctx, source, list)
	return &list
}

func inflateKeyedCustomField(ctx context.Context, source map[string]interface{}) *greenhouse.KeyedCustomField {
	var item greenhouse.KeyedCustomField
	convertType(ctx, source, item)
	return &item
}

func flattenKeyedCustomFields(ctx context.Context, list *map[string]greenhouse.KeyedCustomField) map[string]interface{} {
	flatMap := make(map[string]interface{})
	for k, v := range *list {
		flatMap[k] = flattenKeyedCustomField(ctx, &v)
	}
	return flatMap
}

func flattenKeyedCustomField(ctx context.Context, item *greenhouse.KeyedCustomField) map[string]interface{} {
	kcf := make(map[string]interface{})
	kcf["name"] = item.Name
	kcf["type"] = item.Type
	kcf["value"] = item.Value
	return kcf
}

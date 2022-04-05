package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateKeyedCustomFields(ctx context.Context, source *map[string]interface{}) (*map[string]greenhouse.KeyedCustomField, diag.Diagnostics) {
	list := make(map[string]greenhouse.KeyedCustomField)
	for k, v := range *source {
		itemMap := v.(map[string]interface{})
		obj, err := inflateKeyedCustomField(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[k] = *obj
	}
	return &list, nil
}

func inflateKeyedCustomField(ctx context.Context, source *map[string]interface{}) (*greenhouse.KeyedCustomField, diag.Diagnostics) {
	var obj greenhouse.KeyedCustomField
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["type"].(string); ok && len(v) > 0 {
		obj.Type = &v
	}
	/* TODO this needs to be made consistent with the client.
	if v, ok := (*source)["value"].(string); ok && len(v) > 0 {
	  obj.Value = v
	}
	*/
	return &obj, nil
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
	if v := item.Name; v != nil {
		kcf["name"] = *v
	}
	if v := item.Type; v != nil {
		kcf["type"] = *v
	}
	if v := item.Value; v != nil {
		kcf["value"] = *v
	}
	return kcf
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCustomFieldOption() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the new custom field option.",
			Required:    true,
		},
		"priority": {
			Type:        schema.TypeInt,
			Description: "Numeric value for ordering the custom field options.",
			Required:    true,
		},
		"external_id": {
			Type:        schema.TypeString,
			Description: "The external_id for the custom field.",
			Optional:    true,
		},
	}
}

func inflateCustomFieldOptions(ctx context.Context, source *[]interface{}) (*[]greenhouse.CustomFieldOption, diag.Diagnostics) {
	list := make([]greenhouse.CustomFieldOption, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateCustomFieldOption(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateCustomFieldOption(ctx context.Context, source *map[string]interface{}) (*greenhouse.CustomFieldOption, diag.Diagnostics) {
	var obj greenhouse.CustomFieldOption
	if v, ok := (*source)["external_id"].(string); ok && len(v) > 0 {
		obj.ExternalId = &v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["priority"].(int); ok {
		obj.Priority = &v
	}
	return &obj, nil
}

func flattenCustomFieldOptions(ctx context.Context, list *[]greenhouse.CustomFieldOption) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenCustomFieldOption(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenCustomFieldOption(ctx context.Context, item *greenhouse.CustomFieldOption) map[string]interface{} {
	option := make(map[string]interface{})
	if v := item.ExternalId; v != nil {
		option["external_id"] = *v
	}
	if v := item.Name; v != nil {
		option["name"] = *v
	}
	if v := item.Priority; v != nil {
		option["priority"] = *v
	}
	return option
}

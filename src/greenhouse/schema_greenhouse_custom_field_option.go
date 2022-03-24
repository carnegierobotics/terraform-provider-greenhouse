package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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
	option["external_id"] = item.ExternalId
	option["name"] = item.Name
	option["priority"] = item.Priority
	return option
}

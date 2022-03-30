package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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
	opening["custom_fields"] = item.CustomFields
	return opening
}

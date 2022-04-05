package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseTranslation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"language": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func flattenTranslations(ctx context.Context, list *[]greenhouse.Translation) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenTranslation(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenTranslation(ctx context.Context, item *greenhouse.Translation) map[string]interface{} {
	translation := make(map[string]interface{})
	if v := item.Language; v != nil {
		translation["language"] = *v
	}
	if v := item.Name; v != nil {
		translation["name"] = *v
	}
	return translation
}

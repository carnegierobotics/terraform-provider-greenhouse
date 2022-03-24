package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScorecardAttribute() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"note": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"rating": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func flattenScorecardAttributes(ctx context.Context, list *[]greenhouse.ScorecardAttribute) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenScorecardAttribute(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenScorecardAttribute(ctx context.Context, item *greenhouse.ScorecardAttribute) map[string]interface{} {
	attributes := make(map[string]interface{})
	attributes["name"] = item.Name
	attributes["note"] = item.Note
	attributes["rating"] = item.Rating
	attributes["type"] = item.Type
	return attributes
}

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
	if v := item.Name; v != nil {
		attributes["name"] = *v
	}
	if v := item.Note; v != nil {
		attributes["note"] = *v
	}
	if v := item.Rating; v != nil {
		attributes["rating"] = *v
	}
	if v := item.Type; v != nil {
		attributes["type"] = *v
	}
	return attributes
}

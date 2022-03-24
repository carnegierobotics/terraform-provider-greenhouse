package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCandidateTags() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"candidate_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func flattenCandidateTags(ctx context.Context, list *[]greenhouse.CandidateTag) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			converted := greenhouse.TypeIdName(item)
			flatList[i] = flattenTypeIdName(ctx, &converted)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

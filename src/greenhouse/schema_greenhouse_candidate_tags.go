package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			flatList[i] = flattenCandidateTag(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenCandidateTag(ctx context.Context, item *greenhouse.CandidateTag) map[string]interface{} {
	tag := make(map[string]interface{})
	converted := greenhouse.TypeIdName(*item)
	for k, v := range flattenTypeIdName(ctx, &converted) {
		tag[k] = v
	}
	return tag
}

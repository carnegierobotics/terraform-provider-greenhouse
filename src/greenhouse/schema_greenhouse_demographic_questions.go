package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDemographicQuestions() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"questions": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDemographicQuestion(),
			},
		},
	}
}

func flattenDemographicQuestions(ctx context.Context, list *[]greenhouse.DemographicQuestion) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenDemographicQuestion(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

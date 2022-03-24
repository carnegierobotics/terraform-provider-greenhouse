package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOffers() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"offers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffer(),
			},
		},
	}
}

func flattenOffers(ctx context.Context, list *[]greenhouse.Offer) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenOffer(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

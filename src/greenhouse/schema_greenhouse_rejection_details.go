package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseRejectionDetails() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"custom_fields": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"keyed_custom_fields": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseKeyedCustomField(),
			},
		},
	}
}

func inflateRejectionDetailsList(ctx context.Context, source interface{}) *[]greenhouse.RejectionDetails {
	var list []greenhouse.RejectionDetails
	convertType(ctx, source, list)
	return &list
}

func inflateRejectionDetails(ctx context.Context, source interface{}) *greenhouse.RejectionDetails {
	var item greenhouse.RejectionDetails
	convertType(ctx, source, item)
	return &item
}

func flattenRejectionDetailsList(ctx context.Context, list *[]greenhouse.RejectionDetails) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenRejectionDetails(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenRejectionDetails(ctx context.Context, item *greenhouse.RejectionDetails) map[string]interface{} {
	details := make(map[string]interface{})
	details["custom_fields"] = item.CustomFields
	details["keyed_custom_fields"] = flattenKeyedCustomFields(ctx, &item.KeyedCustomFields)
	return details
}

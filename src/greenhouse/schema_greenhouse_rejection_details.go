package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateRejectionDetailsList(ctx context.Context, source *[]interface{}) (*[]greenhouse.RejectionDetails, diag.Diagnostics) {
	list := make([]greenhouse.RejectionDetails, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateRejectionDetails(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateRejectionDetails(ctx context.Context, source *map[string]interface{}) (*greenhouse.RejectionDetails, diag.Diagnostics) {
	var obj greenhouse.RejectionDetails
	if v, ok := (*source)["custom_fields"].(map[string]string); ok && len(v) > 0 {
		obj.CustomFields = v
	}
	if v, ok := (*source)["keyed_custom_fields"].(map[string]interface{}); ok && len(v) > 0 {
		list, err := inflateKeyedCustomFields(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.KeyedCustomFields = *list
	}
	return &obj, nil
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
	if v := item.CustomFields; len(v) > 0 {
		details["custom_fields"] = v
	}
	if v := item.KeyedCustomFields; len(v) > 0 {
		details["keyed_custom_fields"] = flattenKeyedCustomFields(ctx, &v)
	}
	return details
}

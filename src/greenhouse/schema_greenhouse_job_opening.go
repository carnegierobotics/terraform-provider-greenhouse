package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJobOpening() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"opening_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"close_reason_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"close_reason": {
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"opened_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"closed_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"application_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}
}

func flattenJobOpenings(ctx context.Context, list *[]greenhouse.JobOpening) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			opening := make(map[string]interface{})
			opening["opening_id"] = item.OpeningId
			opening["status"] = item.Status
			opening["opened_at"] = item.OpenedAt
			opening["closed_at"] = item.ClosedAt
			opening["application_id"] = item.ApplicationId
			convertedCloseReason := greenhouse.TypeIdName(item.CloseReason)
			opening["close_reason"] = flattenTypeIdName(ctx, &convertedCloseReason)
			opening["custom_fields"] = item.CustomFields
			flatList[i] = opening
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenCloseReason(ctx context.Context, item *greenhouse.CloseReason) map[string]interface{} {
	flatItem := make(map[string]interface{})
	if item.Name != "" {
		flatItem["name"] = item.Name
	}
	return flatItem
}

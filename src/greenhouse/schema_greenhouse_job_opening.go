package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJobOpening() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"close_reason": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"close_reason_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"closed_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"opened_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"opening_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateJobOpenings(ctx context.Context, source *[]interface{}) (*[]greenhouse.JobOpening, diag.Diagnostics) {
	list := make([]greenhouse.JobOpening, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateJobOpening(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateJobOpening(ctx context.Context, source *map[string]interface{}) (*greenhouse.JobOpening, diag.Diagnostics) {
	var obj greenhouse.JobOpening
	if v, ok := (*source)["application_id"].(int); ok {
		obj.ApplicationId = &v
	}
	if v, ok := (*source)["close_reason"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateCloseReasons(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.CloseReason = &(*list)[0]
	}
	if v, ok := (*source)["closed_at"].(string); ok && len(v) > 0 {
		obj.ClosedAt = &v
	}
	if v, ok := (*source)["custom_fields"].(map[string]string); ok && len(v) > 0 {
		obj.CustomFields = v
	}
	if v, ok := (*source)["opened_at"].(string); ok && len(v) > 0 {
		obj.OpenedAt = &v
	}
	if v, ok := (*source)["opening_id"].(string); ok && len(v) > 0 {
		obj.OpeningId = &v
	}
	if v, ok := (*source)["status"].(string); ok && len(v) > 0 {
		obj.Status = &v
	}
	return &obj, nil
}

func flattenJobOpenings(ctx context.Context, list *[]greenhouse.JobOpening) []interface{} {
	tflog.Debug(ctx, "Flattening job opening list", "opening list", fmt.Sprintf("%+v", list))
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenJobOpening(ctx, &item)
		}
		tflog.Debug(ctx, "Flattened job opening list", "opening list", fmt.Sprintf("%+v", flatList))
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenJobOpening(ctx context.Context, item *greenhouse.JobOpening) map[string]interface{} {
	opening := make(map[string]interface{})
	if v := item.OpeningId; v != nil {
		opening["opening_id"] = *v
	}
	if v := item.Status; v != nil {
		opening["status"] = *v
	}
	if v := item.OpenedAt; v != nil {
		opening["opened_at"] = *v
	}
	if v := item.ClosedAt; v != nil {
		opening["closed_at"] = *v
	}
	if v := item.ApplicationId; v != nil {
		opening["application_id"] = *v
	}
	if v := item.CloseReason; v != nil {
		convertedCloseReason := greenhouse.TypeIdName(*v)
		tflog.Debug(ctx, "Converted close reason", "reason", fmt.Sprintf("%+v", convertedCloseReason))
		opening["close_reason"] = flattenTypeIdName(ctx, &convertedCloseReason)
	} else {
		opening["close_reason"] = nil
	}
	if v := item.CustomFields; len(v) > 0 {
		opening["custom_fields"] = v
	}
	return opening
}

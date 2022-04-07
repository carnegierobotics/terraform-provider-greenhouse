package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseActivity() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"body": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"user": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
	}
}

func inflateActivities(ctx context.Context, source *[]interface{}) (*[]greenhouse.Activity, diag.Diagnostics) {
	list := make([]greenhouse.Activity, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateActivity(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateActivity(ctx context.Context, source *map[string]interface{}) (*greenhouse.Activity, diag.Diagnostics) {
	var obj greenhouse.Activity
	if v, ok := (*source)["body"].(string); ok && len(v) > 0 {
		obj.Body = &v
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
	}
	if v, ok := (*source)["subject"].(string); ok && len(v) > 0 {
		obj.Subject = &v
	}
	if v, ok := (*source)["user"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateUsers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.User = &(*list)[0]
	}
	return &obj, nil
}

func flattenActivities(ctx context.Context, list *[]greenhouse.Activity) []interface{} {
	if list != nil {
		tflog.Trace(ctx, "Flattening activities.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, activity := range *list {
			activity := flattenActivity(ctx, &activity)
			flatList[i] = activity
		}
		tflog.Trace(ctx, "Finished flattening activities.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenActivity(ctx context.Context, item *greenhouse.Activity) map[string]interface{} {
	tflog.Trace(ctx, "Flattening one activity.")
	activity := make(map[string]interface{})
	if v := item.Body; v != nil {
		activity["body"] = *v
	}
	if v := item.CreatedAt; v != nil {
		activity["created_at"] = *v
	}
	if v := item.Subject; v != nil {
		activity["subject"] = *v
	}
	if v := item.User; v != nil {
		activity["user"] = flattenUsersBasics(ctx, &[]greenhouse.User{*v})
	}
	tflog.Trace(ctx, "Finished flattening activity.")
	return activity
}

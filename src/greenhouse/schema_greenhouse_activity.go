package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
	}
}

func flattenActivities(ctx context.Context, list *[]greenhouse.Activity) []interface{} {
	if list != nil {
		tflog.Debug(ctx, "Flattening activities.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, activity := range *list {
			activity := flattenActivity(ctx, &activity)
			flatList[i] = activity
		}
		tflog.Debug(ctx, "Finished flattening activities.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenActivity(ctx context.Context, item *greenhouse.Activity) map[string]interface{} {
	tflog.Debug(ctx, "Flattening one activity.")
	activity := make(map[string]interface{})
	activity["body"] = item.Body
	activity["created_at"] = item.CreatedAt
	activity["subject"] = item.Subject
	activity["user"] = flattenUsersBasics(ctx, &[]greenhouse.User{*item.User})
	tflog.Debug(ctx, "Finished flattening activity.")
	return activity
}

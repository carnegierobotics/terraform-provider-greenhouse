package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
	}
}

func flattenActivities(ctx context.Context, list *[]greenhouse.Activity) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, activity := range *list {
			activity, _ := flattenActivity(ctx, &activity)
			flatList[i] = activity
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenActivity(ctx context.Context, item *greenhouse.Activity) (map[string]interface{}, error) {
	activity := make(map[string]interface{})
	activity["body"] = item.Body
	activity["created_at"] = item.CreatedAt
	activity["subject"] = item.Subject
	activity["user"] = flattenUserBasics(ctx, &item.User)
	return activity, nil
}

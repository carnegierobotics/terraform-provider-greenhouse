package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScheduledInterviewDate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"date_time": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func flattenScheduledInterviewDates(ctx context.Context, list *[]greenhouse.ScheduledInterviewDate) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenScheduledInterviewDate(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenScheduledInterviewDate(ctx context.Context, item *greenhouse.ScheduledInterviewDate) map[string]interface{} {
	date := make(map[string]interface{})
	date["date"] = item.Date
	date["date_time"] = item.DateTime
	return date
}

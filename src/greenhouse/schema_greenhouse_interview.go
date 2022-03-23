package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterview() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_interviewer_users": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterviewer(),
			},
		},
		"estimated_minutes": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"interview_kit": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterviewKit(),
			},
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"schedulable": {
			Type:     schema.TypeBool,
			Optional: true,
		},
	}
}

func flattenInterviews(ctx context.Context, list *[]greenhouse.Interview) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenInterview(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenInterview(ctx context.Context, item *greenhouse.Interview) map[string]interface{} {
	interview := make(map[string]interface{})
	interview["default_interviewer_users"] = flattenInterviewers(ctx, &item.DefaultInterviewerUsers)
	interview["estimated_minutes"] = item.EstimatedMinutes
	interview["interview_kit"] = flattenInterviewKit(ctx, item.InterviewKit)[0]
	interview["name"] = item.Name
	interview["schedulable"] = item.Schedulable
	return interview
}

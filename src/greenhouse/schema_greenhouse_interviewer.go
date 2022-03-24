package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterviewer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"response_status": {
			Type:     schema.TypeString,
			Required: true,
		},
		"scorecard_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func flattenInterviewers(ctx context.Context, list *[]greenhouse.Interviewer) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenInterviewer(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenInterviewer(ctx context.Context, item *greenhouse.Interviewer) map[string]interface{} {
	interviewer := make(map[string]interface{})
	interviewer["email"] = item.Email
	interviewer["employee_id"] = item.EmployeeId
	interviewer["first_name"] = item.FirstName
	interviewer["last_name"] = item.LastName
	interviewer["name"] = item.Name
	interviewer["response_status"] = item.ResponseStatus
	interviewer["scorecard_id"] = item.ScorecardId
	interviewer["user_id"] = item.UserId
	return interviewer
}

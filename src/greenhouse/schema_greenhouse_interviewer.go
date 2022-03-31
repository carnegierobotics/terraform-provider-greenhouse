package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateInterviewers(ctx context.Context, source *[]interface{}) (*[]greenhouse.Interviewer, diag.Diagnostics) {
	list := make([]greenhouse.Interviewer, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateInterviewer(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateInterviewer(ctx context.Context, source *map[string]interface{}) (*greenhouse.Interviewer, diag.Diagnostics) {
	var obj greenhouse.Interviewer
	if v, ok := (*source)["email"].(string); ok && len(v) > 0 {
		obj.Email = v
	}
	if v, ok := (*source)["employee_id"].(string); ok && len(v) > 0 {
		obj.EmployeeId = v
	}
	if v, ok := (*source)["first_name"].(string); ok && len(v) > 0 {
		obj.FirstName = v
	}
	if v, ok := (*source)["last_name"].(string); ok && len(v) > 0 {
		obj.LastName = v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = v
	}
	if v, ok := (*source)["response_status"].(string); ok && len(v) > 0 {
		obj.ResponseStatus = v
	}
	if v, ok := (*source)["scorecard_id"].(int); ok {
		obj.ScorecardId = v
	}
	if v, ok := (*source)["user_id"].(int); ok {
		obj.UserId = v
	}
	return &obj, nil
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

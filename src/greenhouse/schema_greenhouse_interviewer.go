/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of terraform-provider-greenhouse.

terraform-provider-greenhouse is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

terraform-provider-greenhouse is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with terraform-provider-greenhouse. If not, see <https://www.gnu.org/licenses/>.
*/
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
		obj.Email = &v
	}
	if v, ok := (*source)["employee_id"].(string); ok && len(v) > 0 {
		obj.EmployeeId = &v
	}
	if v, ok := (*source)["first_name"].(string); ok && len(v) > 0 {
		obj.FirstName = &v
	}
	if v, ok := (*source)["last_name"].(string); ok && len(v) > 0 {
		obj.LastName = &v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["response_status"].(string); ok && len(v) > 0 {
		obj.ResponseStatus = &v
	}
	if v, ok := (*source)["scorecard_id"].(int); ok {
		obj.ScorecardId = &v
	}
	if v, ok := (*source)["user_id"].(int); ok {
		obj.UserId = &v
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
	if v := item.Email; v != nil {
		interviewer["email"] = *v
	}
	if v := item.EmployeeId; v != nil {
		interviewer["employee_id"] = *v
	}
	if v := item.FirstName; v != nil {
		interviewer["first_name"] = *v
	}
	if v := item.LastName; v != nil {
		interviewer["last_name"] = *v
	}
	if v := item.Name; v != nil {
		interviewer["name"] = *v
	}
	if v := item.ResponseStatus; v != nil {
		interviewer["response_status"] = *v
	}
	if v := item.ScorecardId; v != nil {
		interviewer["scorecard_id"] = *v
	}
	if v := item.UserId; v != nil {
		interviewer["user_id"] = *v
	}
	return interviewer
}

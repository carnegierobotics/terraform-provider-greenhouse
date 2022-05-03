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

func schemaGreenhouseInterview() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"default_interviewer_users": {
			Type:     schema.TypeList,
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
			Type:     schema.TypeList,
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

func inflateInterviews(ctx context.Context, source *[]interface{}) (*[]greenhouse.Interview, diag.Diagnostics) {
	list := make([]greenhouse.Interview, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateInterview(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateInterview(ctx context.Context, source *map[string]interface{}) (*greenhouse.Interview, diag.Diagnostics) {
	var obj greenhouse.Interview
	if v, ok := (*source)["default_interviewer_users"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.DefaultInterviewerUsers = *list
	}
	if v, ok := (*source)["estimated_minutes"].(int); ok {
		obj.EstimatedMinutes = &v
	}
	if v, ok := (*source)["interview_kit"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewKits(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.InterviewKit = &(*list)[0]
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["schedulable"].(bool); ok {
		obj.Schedulable = &v
	}
	return &obj, nil
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
	if v := item.DefaultInterviewerUsers; len(v) > 0 {
		interview["default_interviewer_users"] = flattenInterviewers(ctx, &v)
	}
	if v := item.EstimatedMinutes; v != nil {
		interview["estimated_minutes"] = *v
	}
	if v := item.InterviewKit; v != nil {
		interview["interview_kit"] = flattenInterviewKit(ctx, v)[0]
	}
	if v := item.Name; v != nil {
		interview["name"] = *v
	}
	if v := item.Schedulable; v != nil {
		interview["schedulable"] = *v
	}
	return interview
}

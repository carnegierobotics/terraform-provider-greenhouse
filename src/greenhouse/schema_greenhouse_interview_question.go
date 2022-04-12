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

func schemaGreenhouseInterviewQuestion() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"question": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func inflateInterviewQuestions(ctx context.Context, source *[]interface{}) (*[]greenhouse.InterviewQuestion, diag.Diagnostics) {
	list := make([]greenhouse.InterviewQuestion, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateInterviewQuestion(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateInterviewQuestion(ctx context.Context, source *map[string]interface{}) (*greenhouse.InterviewQuestion, diag.Diagnostics) {
	var obj greenhouse.InterviewQuestion
	if v, ok := (*source)["question"].(string); ok && len(v) > 0 {
		obj.Question = &v
	}
	return &obj, nil
}

func flattenInterviewQuestions(ctx context.Context, list *[]greenhouse.InterviewQuestion) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenInterviewQuestion(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenInterviewQuestion(ctx context.Context, item *greenhouse.InterviewQuestion) map[string]interface{} {
	question := make(map[string]interface{})
	if v := item.Question; v != nil {
		question["question"] = *v
	}
	return question
}

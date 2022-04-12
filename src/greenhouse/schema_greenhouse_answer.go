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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseAnswer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"answer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"question": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateAnswers(ctx context.Context, source *[]interface{}) (*[]greenhouse.Answer, diag.Diagnostics) {
	list := make([]greenhouse.Answer, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateAnswer(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateAnswer(ctx context.Context, source *map[string]interface{}) (*greenhouse.Answer, diag.Diagnostics) {
	var obj greenhouse.Answer
	if v, ok := (*source)["answer"].(string); ok && len(v) > 0 {
		obj.Answer = &v
	}
	if v, ok := (*source)["question"].(string); ok && len(v) > 0 {
		obj.Question = &v
	}
	return &obj, nil
}

func flattenAnswers(ctx context.Context, list *[]greenhouse.Answer) []interface{} {
	if list != nil {
		tflog.Trace(ctx, "Flattening answers.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenAnswer(ctx, &item)
		}
		tflog.Trace(ctx, "Finished flattening answers.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenAnswer(ctx context.Context, item *greenhouse.Answer) map[string]interface{} {
	tflog.Trace(ctx, "Flattening one answer.")
	answer := make(map[string]interface{})
	if v := item.Question; v != nil {
		answer["question"] = *v
	}
	if v := item.Answer; v != nil {
		answer["answer"] = *v
	}
	tflog.Trace(ctx, "Finished flattening answer.")
	return answer
}

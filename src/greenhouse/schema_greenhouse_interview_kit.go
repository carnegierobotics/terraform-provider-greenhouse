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

func schemaGreenhouseInterviewKit() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Type:     schema.TypeString,
			Required: true,
		},
		"questions": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterviewQuestion(),
			},
		},
	}
}

func inflateInterviewKits(ctx context.Context, source *[]interface{}) (*[]greenhouse.InterviewKit, diag.Diagnostics) {
	list := make([]greenhouse.InterviewKit, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateInterviewKit(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateInterviewKit(ctx context.Context, source *map[string]interface{}) (*greenhouse.InterviewKit, diag.Diagnostics) {
	var obj greenhouse.InterviewKit
	if v, ok := (*source)["content"].(string); ok && len(v) > 0 {
		obj.Content = &v
	}
	if v, ok := (*source)["questions"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewQuestions(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Questions = *list
	}
	return &obj, nil
}

func flattenInterviewKit(ctx context.Context, item *greenhouse.InterviewKit) []interface{} {
	kit := make([]interface{}, 1, 1)
	oneKit := make(map[string]interface{})
	if v := item.Content; v != nil {
		oneKit["content"] = *v
	}
	if v := item.Questions; len(v) > 0 {
		oneKit["questions"] = flattenInterviewQuestions(ctx, &v)
	}
	kit[0] = oneKit
	return kit
}

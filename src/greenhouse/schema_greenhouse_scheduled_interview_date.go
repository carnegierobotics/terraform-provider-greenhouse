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

func inflateScheduledInterviewDates(ctx context.Context, source *[]interface{}) (*[]greenhouse.ScheduledInterviewDate, diag.Diagnostics) {
	list := make([]greenhouse.ScheduledInterviewDate, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateScheduledInterviewDate(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateScheduledInterviewDate(ctx context.Context, source *map[string]interface{}) (*greenhouse.ScheduledInterviewDate, diag.Diagnostics) {
	var obj greenhouse.ScheduledInterviewDate
	if v, ok := (*source)["date"].(string); ok && len(v) > 0 {
		obj.Date = &v
	}
	if v, ok := (*source)["date_time"].(string); ok && len(v) > 0 {
		obj.DateTime = &v
	}
	return &obj, nil
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
	if v := item.Date; v != nil {
		date["date"] = *v
	}
	if v := item.DateTime; v != nil {
		date["date_time"] = *v
	}
	return date
}

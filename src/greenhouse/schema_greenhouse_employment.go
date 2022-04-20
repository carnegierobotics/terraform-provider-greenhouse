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

func schemaGreenhouseEmployment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"company_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"title": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateEmployments(ctx context.Context, source *[]interface{}) (*[]greenhouse.Employment, diag.Diagnostics) {
	list := make([]greenhouse.Employment, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		item, diagErr := inflateEmployment(ctx, &itemMap)
		if diagErr != nil {
			return nil, diagErr
		}
		list[i] = *item
	}
	return &list, nil
}

func inflateEmployment(ctx context.Context, source *map[string]interface{}) (*greenhouse.Employment, diag.Diagnostics) {
	var obj greenhouse.Employment
	if v, ok := (*source)["company_name"].(string); ok && len(v) > 0 {
		obj.CompanyName = &v
	}
	if v, ok := (*source)["end_date"].(string); ok && len(v) > 0 {
		obj.EndDate = &v
	}
	if v, ok := (*source)["start_date"].(string); ok && len(v) > 0 {
		obj.StartDate = &v
	}
	if v, ok := (*source)["title"].(string); ok && len(v) > 0 {
		obj.Title = &v
	}
	return &obj, nil
}

func flattenEmployments(ctx context.Context, list *[]greenhouse.Employment) []interface{} {
  if list != nil {
    flatList := make([]interface{}, len(*list), len(*list))
    for i, item := range *list {
      flatList[i] = flattenEmployment(ctx, &item)
    }
    return flatList
  }
  return make([]interface{}, 0, 0)
}

func flattenEmployment(ctx context.Context, item *greenhouse.Employment) map[string]interface{} {
  employment := make(map[string]interface{})
  if v := item.CompanyName; v != nil {
    employment["company_name"] = *v
  }
  if v := item.EndDate; v != nil {
    employment["end_date"] = *v
  }
  if v := item.StartDate; v != nil {
    employment["start_date"] = *v
  }
  if v := item.Title; v != nil {
    employment["title"] = *v
  }
  return employment
}

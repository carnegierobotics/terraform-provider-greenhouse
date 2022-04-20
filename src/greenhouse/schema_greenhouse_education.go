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

func schemaGreenhouseEducation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"degree": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"degree_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"discipline": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"discipline_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_month": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end_year": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"school_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"school_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_month": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_year": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateEducations(ctx context.Context, source *[]interface{}) (*[]greenhouse.Education, diag.Diagnostics) {
	list := make([]greenhouse.Education, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateEducation(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateEducation(ctx context.Context, source *map[string]interface{}) (*greenhouse.Education, diag.Diagnostics) {
	var obj greenhouse.Education
	if v, ok := (*source)["degree"].(string); ok && len(v) > 0 {
		obj.Degree = &v
	}
	if v, ok := (*source)["degree_id"].(int); ok {
		obj.DegreeId = &v
	}
	if v, ok := (*source)["discipline"].(string); ok && len(v) > 0 {
		obj.Discipline = &v
	}
	if v, ok := (*source)["discipline_id"].(int); ok {
		obj.DisciplineId = &v
	}
	if v, ok := (*source)["end_date"].(string); ok && len(v) > 0 {
		obj.EndDate = &v
	}
	if v, ok := (*source)["end_month"].(string); ok && len(v) > 0 {
		obj.EndMonth = &v
	}
	if v, ok := (*source)["end_year"].(string); ok && len(v) > 0 {
		obj.EndYear = &v
	}
	if v, ok := (*source)["school_id"].(int); ok {
		obj.SchoolId = &v
	}
	if v, ok := (*source)["school_name"].(string); ok && len(v) > 0 {
		obj.SchoolName = &v
	}
	if v, ok := (*source)["start_date"].(string); ok && len(v) > 0 {
		obj.StartDate = &v
	}
	if v, ok := (*source)["start_month"].(string); ok && len(v) > 0 {
		obj.StartMonth = &v
	}
	if v, ok := (*source)["start_year"].(string); ok && len(v) > 0 {
		obj.StartYear = &v
	}
	return &obj, nil
}

func flattenEducations(ctx context.Context, list *[]greenhouse.Education) []interface{} {
  if list != nil {
    flatList := make([]interface{}, len(*list), len(*list))
    for i, item := range *list {
      flatList[i] = flattenEducation(ctx, &item)
    }
    return flatList
  }
  return make([]interface{}, 0, 0)
}

func flattenEducation(ctx context.Context, item *greenhouse.Education) map[string]interface{} {
  education := make(map[string]interface{})
  if v := item.Degree; v != nil {
    education["degree"] = *v
  }
  if v := item.Discipline; v != nil {
    education["discipline"] = *v
  }
  if v := item.EndDate; v != nil {
    education["end_date"] = *v
  }
  if v := item.SchoolName; v != nil {
    education["school_name"] = *v
  }
  if v := item.StartDate; v != nil {
    education["start_date"] = *v
  }
  return education
}

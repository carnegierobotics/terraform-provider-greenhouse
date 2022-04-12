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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScorecardAttribute() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"note": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"rating": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func flattenScorecardAttributes(ctx context.Context, list *[]greenhouse.ScorecardAttribute) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenScorecardAttribute(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenScorecardAttribute(ctx context.Context, item *greenhouse.ScorecardAttribute) map[string]interface{} {
	attributes := make(map[string]interface{})
	if v := item.Name; v != nil {
		attributes["name"] = *v
	}
	if v := item.Note; v != nil {
		attributes["note"] = *v
	}
	if v := item.Rating; v != nil {
		attributes["rating"] = *v
	}
	if v := item.Type; v != nil {
		attributes["type"] = *v
	}
	return attributes
}

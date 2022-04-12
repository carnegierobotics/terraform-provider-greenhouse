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
	"strconv"
)

func schemaGreenhouseTypeIdName() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateTypesIdName(ctx context.Context, source *[]interface{}) (*[]greenhouse.TypeIdName, diag.Diagnostics) {
	list := make([]greenhouse.TypeIdName, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateTypeIdName(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateTypeIdName(ctx context.Context, source *map[string]interface{}) (*greenhouse.TypeIdName, diag.Diagnostics) {
	var item greenhouse.TypeIdName
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		item.Name = &v
	}
	return &item, nil
}

func flattenTypeIdName(ctx context.Context, item *greenhouse.TypeIdName) map[string]interface{} {
	newItem := make(map[string]interface{})
	if v := item.Id; v != nil {
		newItem["id"] = strconv.Itoa(*v)
	}
	if v := item.Name; v != nil {
		newItem["name"] = *v
	}
	return newItem
}

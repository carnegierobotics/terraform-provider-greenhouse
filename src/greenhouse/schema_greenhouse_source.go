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

func schemaGreenhouseSource() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func inflateSource(ctx context.Context, source *map[string]interface{}) (*greenhouse.Source, diag.Diagnostics) {
	var obj greenhouse.Source
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["public_name"].(string); ok && len(v) > 0 {
		obj.PublicName = &v
	}
	if v, ok := (*source)["type"].([]interface{}); ok && len(v) > 0 {
		item, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Type = &(*item)[0]
	}
	return &obj, nil
}

func flattenSource(ctx context.Context, item *greenhouse.Source) map[string]interface{} {
	source := make(map[string]interface{})
	if v := item.PublicName; v != nil {
		source["public_name"] = v
	}
	if v := item.Type; v != nil {
		source["type"] = flattenTypeIdName(ctx, v)
	}
	return source
}

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

func schemaGreenhouseRejectionReason() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"include_defaults": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"per_page": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"type": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func inflateRejectionReasons(ctx context.Context, source *[]interface{}) (*[]greenhouse.RejectionReason, diag.Diagnostics) {
	list := make([]greenhouse.RejectionReason, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateRejectionReason(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateRejectionReason(ctx context.Context, source *map[string]interface{}) (*greenhouse.RejectionReason, diag.Diagnostics) {
	var obj greenhouse.RejectionReason
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["type"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Type = &(*list)[0]
	}
	return &obj, nil
}

func flattenRejectionReason(ctx context.Context, item *greenhouse.RejectionReason) map[string]interface{} {
	reason := make(map[string]interface{})
	if v := item.Name; v != nil {
		reason["name"] = item.Name
	}
	if v := item.Type; v != nil {
		reason["type"] = flattenTypeIdName(ctx, v)
	}
	return reason
}

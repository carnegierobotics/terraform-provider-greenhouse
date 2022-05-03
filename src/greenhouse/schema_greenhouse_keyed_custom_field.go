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

func schemaGreenhouseKeyedCustomField() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"value": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func inflateKeyedCustomFields(ctx context.Context, source *map[string]interface{}) (*map[string]greenhouse.KeyedCustomField, diag.Diagnostics) {
	list := make(map[string]greenhouse.KeyedCustomField)
	for k, v := range *source {
		itemMap := v.(map[string]interface{})
		obj, err := inflateKeyedCustomField(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[k] = *obj
	}
	return &list, nil
}

func inflateKeyedCustomField(ctx context.Context, source *map[string]interface{}) (*greenhouse.KeyedCustomField, diag.Diagnostics) {
	var obj greenhouse.KeyedCustomField
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["type"].(string); ok && len(v) > 0 {
		obj.Type = &v
	}
	/* TODO this needs to be made consistent with the client.
	if v, ok := (*source)["value"].(string); ok && len(v) > 0 {
	  obj.Value = v
	}
	*/
	return &obj, nil
}

func flattenKeyedCustomFields(ctx context.Context, list *map[string]greenhouse.KeyedCustomField) map[string]interface{} {
	flatMap := make(map[string]interface{})
	for k, v := range *list {
		flatMap[k] = flattenKeyedCustomField(ctx, &v)
	}
	return flatMap
}

func flattenKeyedCustomField(ctx context.Context, item *greenhouse.KeyedCustomField) map[string]interface{} {
	kcf := make(map[string]interface{})
	if v := item.Name; v != nil {
		kcf["name"] = *v
	}
	if v := item.Type; v != nil {
		kcf["type"] = *v
	}
	if v := item.Value; v != nil {
		kcf["value"] = *v
	}
	return kcf
}

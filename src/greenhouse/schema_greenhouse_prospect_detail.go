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

func schemaGreenhouseProspectDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prospect_owner": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"prospect_pool": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"prospect_stage": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func inflateProspectDetails(ctx context.Context, source *[]interface{}) (*[]greenhouse.ProspectDetail, diag.Diagnostics) {
	list := make([]greenhouse.ProspectDetail, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateProspectDetail(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateProspectDetail(ctx context.Context, source *map[string]interface{}) (*greenhouse.ProspectDetail, diag.Diagnostics) {
	var obj greenhouse.ProspectDetail
	if v, ok := (*source)["prospect_owner"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.ProspectOwner = &(*list)[0]
	}
	if v, ok := (*source)["prospect_pool"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.ProspectPool = &(*list)[0]
	}
	if v, ok := (*source)["prospect_stage"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateTypesIdName(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.ProspectStage = &(*list)[0]
	}
	return &obj, nil
}

func flattenProspectDetail(ctx context.Context, item *greenhouse.ProspectDetail) map[string]interface{} {
	detail := make(map[string]interface{})
	if v := item.ProspectOwner; v != nil {
		detail["prospect_owner"] = flattenTypeIdName(ctx, v)
	}
	if v := item.ProspectPool; v != nil {
		detail["prospect_pool"] = flattenTypeIdName(ctx, v)
	}
	if v := item.ProspectStage; v != nil {
		detail["prospect_stage"] = flattenTypeIdName(ctx, item.ProspectStage)
	}
	return detail
}

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
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOffice() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"child_ids": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"location": {
			/* This is how it should be, but state is not ready.
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseLocation(),
			},
			*/
			//So this is how it is...
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"location_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"primary_contact_user_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}
}

func inflateOffices(ctx context.Context, source *[]interface{}) (*[]greenhouse.Office, diag.Diagnostics) {
	list := make([]greenhouse.Office, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateOffice(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateOffice(ctx context.Context, source *map[string]interface{}) (*greenhouse.Office, diag.Diagnostics) {
	var obj greenhouse.Office
	if v, ok := (*source)["child_ids"].([]interface{}); ok && len(v) > 0 {
		obj.ChildIds = *sliceItoSliceD(&v)
	}
	if v, ok := (*source)["child_office_external_ids"].([]interface{}); ok && len(v) > 0 {
		obj.ChildOfficeExternalIds = *sliceItoSliceA(&v)
	}
	if v, ok := (*source)["external_id"].(string); ok && len(v) > 0 {
		obj.ExternalId = &v
	}
	if v, ok := (*source)["location"].(map[string]interface{}); ok && len(v) > 0 {
		item, err := inflateLocation(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Location = item
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["parent_id"].(int); ok {
		obj.ParentId = &v
	}
	if v, ok := (*source)["parent_office_external_id"].(string); ok && len(v) > 0 {
		obj.ParentOfficeExternalId = &v
	}
	if v, ok := (*source)["primary_contact_user_id"].(int); ok {
		obj.PrimaryContactUserId = &v
	}
	return &obj, nil
}

func flattenOffices(ctx context.Context, list *[]greenhouse.Office) []interface{} {
	tflog.Trace(ctx, "Flattening office list", "officelist", fmt.Sprintf("%+v", list))
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenOffice(ctx, &item)
		}
		tflog.Trace(ctx, "Flattened office list", "officelist", fmt.Sprintf("%+v", flatList))
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenOffice(ctx context.Context, item *greenhouse.Office) map[string]interface{} {
	tflog.Trace(ctx, "Flattening office", "office", fmt.Sprintf("%+v", item))
	office := make(map[string]interface{})
	if v := item.Name; v != nil {
		office["name"] = *v
	}
	if v := item.Location; v != nil {
		office["location"] = flattenLocation(ctx, v)[0]
	}
	if v := item.PrimaryContactUserId; v != nil {
		office["primary_contact_user_id"] = *v
	}
	if v := item.ParentId; v != nil {
		office["parent_id"] = *v
	}
	if v := item.ChildIds; len(v) > 0 {
		office["child_ids"] = v
	}
	tflog.Trace(ctx, "Flattened office", "office", fmt.Sprintf("%+v", office))
	return office
}

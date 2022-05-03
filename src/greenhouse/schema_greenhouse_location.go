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

func schemaGreenhouseLocation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateLocations(ctx context.Context, source *[]interface{}) (*[]greenhouse.Location, diag.Diagnostics) {
	list := make([]greenhouse.Location, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateLocation(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateLocation(ctx context.Context, source *map[string]interface{}) (*greenhouse.Location, diag.Diagnostics) {
	var obj greenhouse.Location
	/*
	  if v, ok := (*source)["address"].(string); ok && len(v) > 0 {
	    obj.Address = v
	  }
	*/
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	return &obj, nil
}

func flattenLocation(ctx context.Context, item *greenhouse.Location) []interface{} {
	tflog.Trace(ctx, "Flattening location", "location", fmt.Sprintf("%+v", item))
	location := make([]interface{}, 1, 1)
	oneLocation := make(map[string]interface{})
	if v := item.Name; v != nil {
		oneLocation["name"] = *v
	}
	location[0] = oneLocation
	tflog.Trace(ctx, "Flattened location", "location", fmt.Sprintf("%+v", location))
	return location
}

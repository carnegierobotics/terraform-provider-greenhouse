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

func schemaGreenhouseApplicationHire() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"close_reason_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"opening_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateApplicationHires(ctx context.Context, source *[]interface{}) (*[]greenhouse.ApplicationHire, diag.Diagnostics) {
	list := make([]greenhouse.ApplicationHire, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApplicationHire(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApplicationHire(ctx context.Context, source *map[string]interface{}) (*greenhouse.ApplicationHire, diag.Diagnostics) {
	var obj greenhouse.ApplicationHire
	if v, ok := (*source)["close_reason_id"].(int); ok {
		obj.CloseReasonId = &v
	}
	if v, ok := (*source)["opening_id"].(int); ok {
		obj.OpeningId = &v
	}
	if v, ok := (*source)["start_date"].(string); ok && len(v) > 0 {
		obj.StartDate = &v
	}
	return &obj, nil
}

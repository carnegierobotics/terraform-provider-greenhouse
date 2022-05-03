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

func schemaGreenhouseApprover() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_addresses": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"employee_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateApprovers(ctx context.Context, source *[]interface{}) (*[]greenhouse.Approver, diag.Diagnostics) {
	list := make([]greenhouse.Approver, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApprover(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApprover(ctx context.Context, source *map[string]interface{}) (*greenhouse.Approver, diag.Diagnostics) {
	var obj greenhouse.Approver
	if v, ok := (*source)["email_addresses"].([]string); ok && len(v) > 0 {
		obj.EmailAddresses = v
	}
	if v, ok := (*source)["employee_id"].(string); ok && len(v) > 0 {
		obj.EmployeeId = &v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	return &obj, nil
}

func flattenApprovers(ctx context.Context, list *[]greenhouse.Approver) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApprover(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApprover(ctx context.Context, item *greenhouse.Approver) map[string]interface{} {
	approver := make(map[string]interface{})
	approver["email_addresses"] = item.EmailAddresses
	if v := item.EmployeeId; v != nil {
		approver["employee_id"] = *v
	}
	if v := item.Name; v != nil {
		approver["name"] = *v
	}
	return approver
}

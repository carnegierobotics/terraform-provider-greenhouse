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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUserBasics() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func flattenUsersBasics(ctx context.Context, list *[]greenhouse.User) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenUserBasics(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenUserBasics(ctx context.Context, item *greenhouse.User) map[string]interface{} {
	tflog.Trace(ctx, "Flattening one user basics.")
	user := make(map[string]interface{})
	if v := item.EmployeeId; v != nil {
		user["employee_id"] = *v
	}
	if v := item.FirstName; v != nil {
		user["first_name"] = *v
	}
	if v := item.LastName; v != nil {
		user["last_name"] = *v
	}
	if v := item.Name; v != nil {
		user["name"] = *v
	}
	tflog.Trace(ctx, "Finished flattening user basics.")
	return user
}

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUserPermission() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"user_role_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func flattenUserPermissions(ctx context.Context, list *[]greenhouse.UserPermission) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenUserPermission(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenUserPermission(ctx context.Context, item *greenhouse.UserPermission) map[string]interface{} {
	permission := make(map[string]interface{})
	if v := item.JobId; v != nil {
		permission["job_id"] = *v
	}
	if v := item.UserRoleId; v != nil {
		permission["user_role_id"] = *v
	}
	return permission
}

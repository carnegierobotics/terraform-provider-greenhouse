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

func schemaGreenhouseFutureJobPermission() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"department_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"external_department_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"external_office_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"office_id": {
			Type:     schema.TypeInt,
			Optional: true,
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

func flattenFutureJobPermissions(ctx context.Context, list *[]greenhouse.FutureJobPermission) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenFutureJobPermission(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenFutureJobPermission(ctx context.Context, item *greenhouse.FutureJobPermission) map[string]interface{} {
	permission := make(map[string]interface{})
	if v := item.DepartmentId; v != nil {
		permission["department_id"] = *v
	}
	if v := item.ExternalDepartmentId; v != nil {
		permission["external_department_id"] = *v
	}
	if v := item.ExternalOfficeId; v != nil {
		permission["external_office_id"] = *v
	}
	if v := item.OfficeId; v != nil {
		permission["office_id"] = *v
	}
	if v := item.UserRoleId; v != nil {
		permission["user_role_id"] = *v
	}
	return permission
}

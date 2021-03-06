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

func schemaGreenhouseHiringMember() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		/*
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		*/
		"employee_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"responsible": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"responsible_for_future_candidates": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"responsible_for_active_candidates": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"responsible_for_inactive_candidates": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func inflateHiringSubteamMember(ctx context.Context, source *map[string]interface{}) (*greenhouse.HiringMember, diag.Diagnostics) {
	var obj greenhouse.HiringMember
	if v, ok := (*source)["employee_id"].(string); ok && len(v) > 0 {
		obj.EmployeeId = &v
	}
	if v, ok := (*source)["first_name"].(string); ok && len(v) > 0 {
		obj.FirstName = &v
	}
	if v, ok := (*source)["last_name"].(string); ok && len(v) > 0 {
		obj.LastName = &v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*source)["responsible"].(bool); ok {
		obj.Responsible = &v
	}
	if v, ok := (*source)["responsible_for_future_candidates"].(bool); ok {
		obj.ResponsibleForFutureCandidates = &v
	}
	if v, ok := (*source)["responsible_for_active_candidates"].(bool); ok {
		obj.ResponsibleForActiveCandidates = &v
	}
	if v, ok := (*source)["responsible_for_inactive_candidates"].(bool); ok {
		obj.ResponsibleForInactiveCandidates = &v
	}
	if v, ok := (*source)["user_id"].(int); ok {
		obj.UserId = &v
	}
	return &obj, nil
}

func flattenHiringSubteamMember(ctx context.Context, item greenhouse.HiringMember) (map[string]interface{}, error) {
	tflog.Trace(ctx, "User data", "user", fmt.Sprintf("%+v", item))
	member := make(map[string]interface{})
	/*
	  member["user_id"] = *item.Id
	  member["active"] = *item.Active
	  member["first_name"] = *item.FirstName
	  member["last_name"] = *item.LastName
	  member["name"] = *item.Name
	  member["responsible"] = *item.Responsible
	  member["employee_id"] = *item.EmployeeId
	*/
	if v := item.Id; v != nil {
		member["user_id"] = *v
	}
	if v := item.Active; v != nil {
		member["active"] = *v
	}
	if v := item.FirstName; v != nil {
		member["first_name"] = *v
	}
	if v := item.LastName; v != nil {
		member["last_name"] = *v
	}
	if v := item.Name; v != nil {
		member["name"] = *v
	}
	if v := item.Responsible; v != nil {
		member["responsible"] = *v
	}
	if v := item.EmployeeId; v != nil {
		member["employee_id"] = *v
	}
	return member, nil
}

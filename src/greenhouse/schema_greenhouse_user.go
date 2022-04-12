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

func schemaGreenhouseUser() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"disable_user": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"disabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"emails": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
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
		"linked_candidate_ids": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"primary_email_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"send_email": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"site_admin": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateUsers(ctx context.Context, source *[]interface{}) (*[]greenhouse.User, diag.Diagnostics) {
	tflog.Trace(ctx, fmt.Sprintf("Inflating users: %+v", source))
	list := make([]greenhouse.User, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateUser(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateUser(ctx context.Context, item *map[string]interface{}) (*greenhouse.User, diag.Diagnostics) {
	var obj greenhouse.User
	if v, ok := (*item)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
	}
	if v, ok := (*item)["disabled"].(bool); ok {
		obj.Disabled = &v
	}
	if v, ok := (*item)["emails"].([]string); ok && len(v) > 0 {
		obj.Emails = v
	}
	if v, ok := (*item)["employee_id"].(string); ok && len(v) > 0 {
		obj.EmployeeId = &v
	}
	if v, ok := (*item)["first_name"].(string); ok && len(v) > 0 {
		obj.FirstName = &v
	}
	if v, ok := (*item)["last_name"].(string); ok && len(v) > 0 {
		obj.LastName = &v
	}
	if v, ok := (*item)["linked_candidate_ids"].([]int); ok && len(v) > 0 {
		obj.LinkedCandidateIds = v
	}
	if v, ok := (*item)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := (*item)["primary_email_address"].(string); ok && len(v) > 0 {
		obj.PrimaryEmail = &v
	}
	if v, ok := (*item)["site_admin"].(bool); ok {
		obj.SiteAdmin = &v
	}
	if v, ok := (*item)["updated_at"].(string); ok {
		obj.UpdatedAt = &v
	}
	return &obj, nil
}

func flattenUser(ctx context.Context, item *greenhouse.User) map[string]interface{} {
	tflog.Trace(ctx, "User item:", fmt.Sprintf("%+v\n", *item))
	user := make(map[string]interface{})
	if v := item.CreatedAt; v != nil {
		user["created_at"] = *v
	}
	if v := item.Disabled; v != nil {
		user["disabled"] = *v
	}
	user["emails"] = item.Emails
	if v := item.EmployeeId; v != nil {
		user["employee_id"] = *v
	}
	if v := item.FirstName; v != nil {
		user["first_name"] = *v
	}
	if v := item.LastName; v != nil {
		user["last_name"] = *v
	}
	user["linked_candidate_ids"] = item.LinkedCandidateIds
	if v := item.Name; v != nil {
		user["name"] = item.Name
	}
	if v := item.PrimaryEmail; v != nil {
		user["primary_email_address"] = *v
	}
	if v := item.SiteAdmin; v != nil {
		user["site_admin"] = *v
	}
	if v := item.UpdatedAt; v != nil {
		user["updated_at"] = *v
	}
	return user
}

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"body": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cc": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"from": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"to": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"user": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
	}
}

func inflateEmails(ctx context.Context, source *[]interface{}) (*[]greenhouse.Email, diag.Diagnostics) {
	list := make([]greenhouse.Email, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateEmail(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateEmail(ctx context.Context, source *map[string]interface{}) (*greenhouse.Email, diag.Diagnostics) {
	var obj greenhouse.Email
	if v, ok := (*source)["body"].(string); ok && len(v) > 0 {
		obj.Body = &v
	}
	if v, ok := (*source)["cc"].(string); ok && len(v) > 0 {
		obj.Cc = &v
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
	}
	if v, ok := (*source)["from"].(string); ok && len(v) > 0 {
		obj.From = &v
	}
	if v, ok := (*source)["subject"].(string); ok && len(v) > 0 {
		obj.Subject = &v
	}
	if v, ok := (*source)["to"].(string); ok && len(v) > 0 {
		obj.To = &v
	}
	if v, ok := (*source)["user"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateUsers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.User = &(*list)[0]
	}
	return &obj, nil
}

func flattenEmails(ctx context.Context, list *[]greenhouse.Email) []interface{} {
	if list != nil {
		tflog.Trace(ctx, "Flattening emails.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, email := range *list {
			email := flattenEmail(ctx, &email)
			flatList[i] = email
		}
		tflog.Trace(ctx, "Finished flattening emails.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenEmail(ctx context.Context, item *greenhouse.Email) map[string]interface{} {
	tflog.Trace(ctx, "Flattening one email.")
	email := make(map[string]interface{})
	if v := item.Body; v != nil {
		email["body"] = *v
	}
	if v := item.Cc; v != nil {
		email["cc"] = String(item.Cc)
	}
	if v := item.CreatedAt; v != nil {
		email["created_at"] = String(item.CreatedAt)
	}
	if v := item.Subject; v != nil {
		email["subject"] = String(item.Subject)
	}
	if v := item.To; v != nil {
		email["to"] = String(item.To)
	}
	if v := item.User; v != nil {
		email["user"] = flattenUsersBasics(ctx, &[]greenhouse.User{*item.User})
	}
	tflog.Trace(ctx, "Finished flattening email.")
	return email
}

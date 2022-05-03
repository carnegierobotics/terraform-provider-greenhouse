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

func schemaGreenhouseApproverGroup() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"approvals_required": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"approvers": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApprover(),
			},
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"offer_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"resolved_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateApproverGroups(ctx context.Context, source *[]interface{}) (*[]greenhouse.ApproverGroup, diag.Diagnostics) {
	list := make([]greenhouse.ApproverGroup, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApproverGroup(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApproverGroup(ctx context.Context, source *map[string]interface{}) (*greenhouse.ApproverGroup, diag.Diagnostics) {
	var obj greenhouse.ApproverGroup
	if v, ok := (*source)["approvals_required"].(int); ok {
		obj.ApprovalsRequired = &v
	}
	if v, ok := (*source)["approvers"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateApprovers(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Approvers = *list
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
	}
	if v, ok := (*source)["job_id"].(int); ok {
		obj.JobId = &v
	}
	if v, ok := (*source)["offer_id"].(int); ok {
		obj.OfferId = &v
	}
	if v, ok := (*source)["priority"].(int); ok {
		obj.Priority = &v
	}
	if v, ok := (*source)["resolved_at"].(string); ok && len(v) > 0 {
		obj.ResolvedAt = &v
	}
	return &obj, nil
}

func flattenApproverGroups(ctx context.Context, list *[]greenhouse.ApproverGroup) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApproverGroup(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApproverGroup(ctx context.Context, item *greenhouse.ApproverGroup) map[string]interface{} {
	approver := make(map[string]interface{})
	if v := item.ApprovalsRequired; v != nil {
		approver["approvals_required"] = *v
	}
	if v := item.Approvers; len(v) > 0 {
		approver["approvers"] = flattenApprovers(ctx, &v)
	}
	if v := item.CreatedAt; v != nil {
		approver["created_at"] = *v
	}
	if v := item.JobId; v != nil {
		approver["job_id"] = *v
	}
	if v := item.OfferId; v != nil {
		approver["offer_id"] = *v
	}
	if v := item.Priority; v != nil {
		approver["priority"] = *v
	}
	if v := item.ResolvedAt; v != nil {
		approver["resolved_at"] = *v
	}
	return approver
}

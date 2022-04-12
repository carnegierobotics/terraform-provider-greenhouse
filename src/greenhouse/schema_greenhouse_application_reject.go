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

func schemaGreenhouseApplicationReject() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"notes": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"rejection_email": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRejectionEmail(),
			},
		},
		"rejection_reason": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func inflateApplicationRejects(ctx context.Context, source *[]interface{}) (*[]greenhouse.ApplicationReject, diag.Diagnostics) {
	list := make([]greenhouse.ApplicationReject, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApplicationReject(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApplicationReject(ctx context.Context, source *map[string]interface{}) (*greenhouse.ApplicationReject, diag.Diagnostics) {
	var obj greenhouse.ApplicationReject
	if v, ok := (*source)["notes"].(string); ok && len(v) > 0 {
		obj.Notes = &v
	}
	if v, ok := (*source)["rejection_email"].([]interface{}); ok && len(v) > 0 {
		item, diagErr := inflateRejectionEmails(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.RejectionEmail = &(*item)[0]
	}
	if v, ok := (*source)["rejection_reason"].(int); ok {
		obj.RejectionReasonId = &v
	}
	return &obj, nil
}

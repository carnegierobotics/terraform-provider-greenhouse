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

func schemaGreenhouseActivityFeed() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activities": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseActivity(),
			},
		},
		"candidate_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"emails": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEmail(),
			},
		},
		"notes": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseNote(),
			},
		},
	}
}

func inflateActivityFeeds(ctx context.Context, source *[]interface{}) (*[]greenhouse.ActivityFeed, diag.Diagnostics) {
	list := make([]greenhouse.ActivityFeed, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateActivityFeed(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateActivityFeed(ctx context.Context, source *map[string]interface{}) (*greenhouse.ActivityFeed, diag.Diagnostics) {
	var obj greenhouse.ActivityFeed
	if v, ok := (*source)["activities"].([]interface{}); ok && len(v) > 0 {
		activities, diagErr := inflateActivities(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Activities = *activities
	}
	if v, ok := (*source)["emails"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateEmails(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Emails = *list
	}
	if v, ok := (*source)["notes"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateNotes(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Notes = *list
	}
	return &obj, nil
}

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

func schemaGreenhouseScheduledInterview() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"end": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseScheduledInterviewDate(),
			},
		},
		"external_event_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"interview_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"interviewers": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterviewer(),
			},
		},
		"location": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"organizer": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUser(),
			},
		},
		"start": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseScheduledInterviewDate(),
			},
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"video_conferencing_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateScheduledInterviews(ctx context.Context, source *[]interface{}) (*[]greenhouse.ScheduledInterview, diag.Diagnostics) {
	list := make([]greenhouse.ScheduledInterview, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateScheduledInterview(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateScheduledInterview(ctx context.Context, source *map[string]interface{}) (*greenhouse.ScheduledInterview, diag.Diagnostics) {
	var obj greenhouse.ScheduledInterview
	if v, ok := (*source)["application_id"].(int); ok {
		obj.ApplicationId = &v
	}
	if v, ok := (*source)["end"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateScheduledInterviewDates(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.End = &(*list)[0]
	}
	if v, ok := (*source)["external_event_id"].(string); ok && len(v) > 0 {
		obj.ExternalEventId = &v
	}
	if v, ok := (*source)["interviewers"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Interviewers = *list
	}
	if v, ok := (*source)["location"].(string); ok && len(v) > 0 {
		obj.Location = &v
	}
	if v, ok := (*source)["organizer"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateUsers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Organizer = &(*list)[0]
	}
	if v, ok := (*source)["start"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateScheduledInterviewDates(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Start = &(*list)[0]
	}
	if v, ok := (*source)["status"].(string); ok && len(v) > 0 {
		obj.Status = &v
	}
	if v, ok := (*source)["video_conferencing_url"].(string); ok && len(v) > 0 {
		obj.VideoConferencingUrl = &v
	}
	return &obj, nil
}

func flattenScheduledInterviews(ctx context.Context, list *[]greenhouse.ScheduledInterview) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenScheduledInterview(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenScheduledInterview(ctx context.Context, item *greenhouse.ScheduledInterview) map[string]interface{} {
	interview := make(map[string]interface{})
	if v := item.ApplicationId; v != nil {
		interview["application_id"] = *v
	}
	if v := item.End; v != nil {
		interview["end"] = flattenScheduledInterviewDate(ctx, v)
	}
	if v := item.ExternalEventId; v != nil {
		interview["external_event_id"] = *v
	}
	if v := item.Interviewers; len(v) > 0 {
		interview["interviewers"] = flattenInterviewers(ctx, &v)
	}
	if v := item.Location; v != nil {
		interview["location"] = *v
	}
	if v := item.Organizer; v != nil {
		interview["organizer"] = flattenUser(ctx, v)
	}
	if v := item.Start; v != nil {
		interview["start"] = flattenScheduledInterviewDate(ctx, v)
	}
	if v := item.Status; v != nil {
		interview["status"] = *v
	}
	if v := item.VideoConferencingUrl; v != nil {
		interview["video_conferencing_url"] = *v
	}
	return interview
}

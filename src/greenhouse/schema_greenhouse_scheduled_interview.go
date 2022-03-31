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
		obj.ApplicationId = v
	}
	if v, ok := (*source)["end"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateScheduledInterviewDates(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.End = &(*list)[0]
	}
	if v, ok := (*source)["external_event_id"].(string); ok && len(v) > 0 {
		obj.ExternalEventId = v
	}
	if v, ok := (*source)["interviewers"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Interviewers = *list
	}
	if v, ok := (*source)["location"].(string); ok && len(v) > 0 {
		obj.Location = v
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
		obj.Status = v
	}
	if v, ok := (*source)["video_conferencing_url"].(string); ok && len(v) > 0 {
		obj.VideoConferencingUrl = v
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
	interview["application_id"] = item.ApplicationId
	interview["end"] = flattenScheduledInterviewDate(ctx, item.End)
	interview["external_event_id"] = item.ExternalEventId
	interview["interviewers"] = flattenInterviewers(ctx, &item.Interviewers)
	interview["location"] = item.Location
	interview["organizer"] = flattenUser(ctx, item.Organizer)
	interview["start"] = flattenScheduledInterviewDate(ctx, item.Start)
	interview["status"] = item.Status
	interview["video_conferencing_url"] = item.VideoConferencingUrl
	return interview
}

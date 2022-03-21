package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScheduledInterview() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end": {
			Type:     schema.TypeSet,
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
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUser(),
			},
		},
		"start": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseScheduledInterviewDate(),
			},
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"video_conferencing_url": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

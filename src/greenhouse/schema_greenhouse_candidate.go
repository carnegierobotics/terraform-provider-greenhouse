package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCandidate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activity_feed_notes": {
			Type:        schema.TypeList,
			Description: "The candidate's activity feed.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseActivityFeed(),
			},
		},
		"addresses": {
			Type:        schema.TypeSet,
			Description: "The candidate's address(es).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"application": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApplication(),
			},
			ConflictsWith: []string{"applications"},
		},
		"application_ids": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Computed: true,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"applications": {
			Type:        schema.TypeSet,
			Description: "Applications the candidate has submitted.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApplication(),
			},
			ConflictsWith: []string{"application"},
		},
		"can_email": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"company": {
			Type:        schema.TypeString,
			Description: "The candidate's company.",
			Optional:    true,
		},
		"coordinator": {
			Type:        schema.TypeList,
			Description: "The candidate's coordinator.",
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRecruiter(),
			},
		},
		"custom_fields": {
			Type:        schema.TypeSet,
			Description: "Custom fields for this candidate.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseCustomField(),
			},
		},
		"educations": {
			Type:        schema.TypeList,
			Description: "The candidate's educational background.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEducation(),
			},
		},
		"email_addresses": {
			Type:        schema.TypeSet,
			Description: "The candidate's email address(es).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"employments": {
			Type:        schema.TypeList,
			Description: "The candidate's employment background.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEmployment(),
			},
		},
		"first_name": {
			Type:        schema.TypeString,
			Description: "The candidate's first name.",
			Required:    true,
		},
		"is_private": {
			Type:        schema.TypeBool,
			Description: "This candidate is private.",
			Optional:    true,
		},
		"is_prospect": {
			Type:        schema.TypeBool,
			Description: "This candidate is a prospect.",
			Required:    true,
		},
		"last_name": {
			Type:        schema.TypeString,
			Description: "The candidate's last name.",
			Required:    true,
		},
		"linked_user_ids": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"phone_numbers": {
			Type:        schema.TypeSet,
			Description: "The candidate's phone number(s).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"recruiter": {
			Type:        schema.TypeList,
			Description: "The candidate's recruiter.",
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRecruiter(),
			},
		},
		"social_media_addresses": {
			Type:        schema.TypeSet,
			Description: "The candidate's social media address(es).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"tags": {
			Type:        schema.TypeSet,
			Description: "Tags for this candidate.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"title": {
			Type:        schema.TypeString,
			Description: "The candidate's title.",
			Optional:    true,
		},
		"website_addresses": {
			Type:        schema.TypeSet,
			Description: "The candidate's website(s).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
	}
}

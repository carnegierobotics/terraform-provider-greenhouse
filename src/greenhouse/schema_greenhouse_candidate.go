package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCandidate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"first_name": {
			Type:     schema.TypeString,
      Description: "The candidate's first name.",
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
      Description: "The candidate's last name.",
			Required: true,
		},
		"company": {
			Type:     schema.TypeString,
      Description: "The candidate's company.",
			Optional: true,
		},
		"title": {
			Type:     schema.TypeString,
      Description: "The candidate's title.",
			Optional: true,
		},
    "is_private": {
      Type:     schema.TypeBool,
      Description: "This candidate is private.",
      Optional: true,
    },
		"phone_numbers": {
			Type:     schema.TypeSet,
      Description: "The candidate's phone number(s).",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseValueType(),
			},
		},
		"addresses": {
			Type:     schema.TypeSet,
      Description: "The candidate's address(es).",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseValueType(),
			},
		},
		"email_addresses": {
			Type:     schema.TypeSet,
      Description: "The candidate's email address(es).",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseValueType(),
			},
		},
		"website_addresses": {
			Type:     schema.TypeSet,
      Description: "The candidate's website(s).",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseValueType(),
			},
		},
		"social_media_addresses": {
			Type:     schema.TypeSet,
      Description: "The candidate's social media address(es).",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseValueType(),
			},
		},
		"educations": {
			Type:     schema.TypeList,
      Description: "The candidate's educational background.",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEducation(),
			},
		},
		"employments": {
			Type:     schema.TypeList,
      Description: "The candidate's employment background.",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEmployment(),
			},
		},
		"tags": {
			Type:     schema.TypeSet,
      Description: "Tags for this candidate.",
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"custom_fields": {
			Type:     schema.TypeSet,
      Description: "Custom fields for this candidate.",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseCustomField(),
			},
		},
		"recruiter": {
			Type:     schema.TypeList,
      Description: "The candidate's recruiter.",
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRecruiter(),
			},
		},
		"coordinator": {
			Type:     schema.TypeList,
      Description: "The candidate's coordinator.",
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRecruiter(),
			},
		},
		"activity_feed_notes": {
			Type:     schema.TypeList,
      Description: "The candidate's activity feed.",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseActivityFeed(),
			},
		},
		"applications": {
			Type:     schema.TypeList,
      Description: "Applications the candidate has submitted.",
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApplication(),
			},
		},
	}
}

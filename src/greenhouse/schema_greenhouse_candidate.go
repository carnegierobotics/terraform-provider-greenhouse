package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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
		},
		"attachments": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseAttachment(),
			},
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
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Computed: true,
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
		"last_activity": {
			Type:     schema.TypeString,
			Computed: true,
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
		"photo_url": {
			Type:     schema.TypeString,
			Optional: true,
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
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
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

func flattenCandidates(ctx context.Context, list *[]greenhouse.Candidate) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenCandidate(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenCandidate(ctx context.Context, item *greenhouse.Candidate) map[string]interface{} {
	candidate := make(map[string]interface{})
	convertedAddresses := []greenhouse.TypeTypeValue(item.Addresses)
	candidate["addresses"] = flattenTypeTypeValues(ctx, &convertedAddresses)
	candidate["application_ids"] = item.ApplicationIds
	candidate["attachments"] = flattenAttachments(ctx, &item.Attachments)
	candidate["company"] = item.Company
	candidate["created_at"] = item.CreatedAt
	candidate["first_name"] = item.FirstName
	candidate["is_private"] = item.IsPrivate
	candidate["last_activity"] = item.LastActivity
	candidate["last_name"] = item.LastName
	convertedPhoneNumbers := []greenhouse.TypeTypeValue(item.PhoneNumbers)
	candidate["phone_numbers"] = flattenTypeTypeValues(ctx, &convertedPhoneNumbers)
	candidate["photo_url"] = item.PhotoUrl
	candidate["title"] = item.Title
	candidate["updated_at"] = item.UpdatedAt
	return candidate
}

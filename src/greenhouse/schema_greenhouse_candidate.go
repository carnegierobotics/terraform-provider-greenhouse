package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			Type:        schema.TypeList,
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
			Type:     schema.TypeList,
			MaxItems: 1,
			Computed: true,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"applications": {
			Type:        schema.TypeList,
			Description: "Applications the candidate has submitted.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApplication(),
			},
		},
		"attachments": {
			Type:     schema.TypeList,
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
      Computed:    true,
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
			Type:        schema.TypeList,
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
			Type:        schema.TypeList,
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
      Computed:    true,
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
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"phone_numbers": {
			Type:        schema.TypeList,
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
			Type:        schema.TypeList,
			Description: "The candidate's social media address(es).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"tags": {
			Type:        schema.TypeList,
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
			Type:        schema.TypeList,
			Description: "The candidate's website(s).",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
	}
}

func inflateCandidates(ctx context.Context, source *[]interface{}) (*[]greenhouse.Candidate, diag.Diagnostics) {
  list := make([]greenhouse.Candidate, len(*source), len(*source))
  for i, item := range *source {
    itemMap := item.(map[string]interface{})
    obj, err := inflateCandidate(ctx, &itemMap)
    if err != nil {
      return nil, err
    }
    list[i] = *obj
  }
  return &list, nil
}

func inflateCandidate(ctx context.Context, source *map[string]interface{}) (*greenhouse.Candidate, diag.Diagnostics) {
  var obj greenhouse.Candidate
  if v, ok := (*source)["activity_feed_notes"].([]interface{}); ok && len(v) > 0 {
    list, diagErr := inflateActivityFeeds(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.ActivityFeedNotes = *list
  }
  if v, ok := (*source)["addresses"].([]interface{}); ok && len(v) > 0 {
    addresses, diagErr := inflateTypeTypeValues(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Addresses = *addresses
  }
  if v, ok := (*source)["application"].([]interface{}); ok && len(v) > 0 {
    item, diagErr := inflateApplications(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Application = &(*item)[0]
  }
  if v, ok := (*source)["application_ids"].([]int); ok && len(v) > 0 {
    obj.ApplicationIds = v
  }
  if v, ok := (*source)["applications"].([]interface{}); ok && len(v) > 0 {
    list, diagErr := inflateApplications(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Applications = *list
  }
  if v, ok := (*source)["attachments"].([]interface{}); ok && len(v) > 0 {
    list, diagErr := inflateAttachments(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Attachments = *list
  }
  if v, ok := (*source)["can_email"].(bool); ok {
    obj.CanEmail = v
  }
  if v, ok := (*source)["company"].(string); ok && len(v) > 0 {
    obj.Company = v
  }
  if v, ok := (*source)["coordinator"].([]interface{}); ok && len(v) > 0 {
    item, err := inflateUser(ctx, &(v[0]))
    if err != nil {
      return nil, err
    }
    obj.Coordinator = item 
  }
  if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
    obj.CreatedAt = v
  }
  if v, ok := (*source)["educations"].([]interface{}); ok && len(v) > 0 {
    list, diagErr := inflateEducations(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Educations = *list
  }
  if v, ok := (*source)["email_addresses"].([]interface{}); ok && len(v) > 0 {
    emails, diagErr := inflateTypeTypeValues(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.EmailAddresses = *emails
  }
  if v, ok := (*source)["employments"].([]interface{}); ok && len(v) > 0 {
    list, diagErr := inflateEmployments(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Employments = *list
  }
  if v, ok := (*source)["first_name"].(string); ok && len(v) > 0 {
    obj.FirstName = v
  }
  if v, ok := (*source)["is_private"].(bool); ok {
    obj.IsPrivate = v
  }
  if v, ok := (*source)["is_prospect"].(bool); ok {
    obj.IsProspect = v
  }
  if v, ok := (*source)["last_activity"].(string); ok && len(v) > 0 {
    obj.LastActivity = v
  }
  if v, ok := (*source)["last_name"].(string); ok && len(v) > 0 {
    obj.LastName = v
  }
  if v, ok := (*source)["linked_user_ids"].([]int); ok && len(v) > 0 {
    obj.LinkedUserIds = v
  }
  if v, ok := (*source)["phone_numbers"].([]interface{}); ok && len(v) > 0 {
    phoneNumbers, diagErr := inflateTypeTypeValues(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.PhoneNumbers = *phoneNumbers
  }
  if v, ok := (*source)["photo_url"].(string); ok && len(v) > 0 {
    obj.PhotoUrl = v
  }
  if v, ok := (*source)["recruiter"].([]interface{}); ok && len(v) > 0 {
    item, diagErr := inflateUser(ctx, &(v[0]))
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Recruiter = item
  }
  if v, ok := (*source)["social_media_addresses"].([]interface{}); ok && len(v) > 0 {
    addresses, diagErr := inflateTypeTypeValues(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.Addresses = *addresses
  }
  if v, ok := (*source)["tags"].([]string); ok && len(v) > 0 {
    obj.Tags = v
  }
  if v, ok := (*source)["title"].(string); ok && len(v) > 0 {
    obj.Title = v
  }
  if v, ok := (*source)["updated_at"].(string); ok && len(v) > 0 {
    obj.UpdatedAt = v
  }
  if v, ok := (*source)["website_addresses"].([]interface{}); ok && len(v) > 0 {
    addresses, diagErr := inflateTypeTypeValues(ctx, &v)
    if diagErr != nil {
      return nil, diagErr
    }
    obj.WebsiteAddresses = *addresses
  }
  return &obj, nil
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

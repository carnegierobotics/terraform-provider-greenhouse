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
		"anonymize": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
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
		"merge": {
			Type:     schema.TypeInt,
			Optional: true,
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
			Computed: true,
		},
		"recruiter": {
			Type:        schema.TypeList,
			Description: "The candidate's recruiter.",
			MaxItems:    1,
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRecruiter(),
			},
		},
		"social_media_addresses": {
			Type:        schema.TypeList,
			Description: "The candidate's social media address(es).",
			Optional:    true,
			Computed:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeTypeValue(),
			},
		},
		"tag_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"tags": {
			Type:        schema.TypeList,
			Description: "Tags for this candidate.",
			Computed:    true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"title": {
			Type:        schema.TypeString,
			Description: "The candidate's title.",
			Optional:    true,
			Computed:    true,
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
	if v, ok := (*source)["application_ids"].([]interface{}); ok && len(v) > 0 {
		ids := *sliceItoSliceD(&v)
		obj.ApplicationIds = ids
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
		obj.CanEmail = &v
	}
	if v, ok := (*source)["company"].(string); ok && len(v) > 0 {
		obj.Company = &v
	}
	if v, ok := (*source)["coordinator"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateUsers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Coordinator = &(*list)[0]
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
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
		obj.FirstName = &v
	}
	if v, ok := (*source)["is_private"].(bool); ok {
		obj.IsPrivate = &v
	}
	if v, ok := (*source)["is_prospect"].(bool); ok {
		obj.IsProspect = &v
	}
	if v, ok := (*source)["last_activity"].(string); ok && len(v) > 0 {
		obj.LastActivity = &v
	}
	if v, ok := (*source)["last_name"].(string); ok && len(v) > 0 {
		obj.LastName = &v
	}
	if v, ok := (*source)["linked_user_ids"].([]interface{}); ok && len(v) > 0 {
		ids := *sliceItoSliceD(&v)
		obj.LinkedUserIds = ids
	}
	if v, ok := (*source)["phone_numbers"].([]interface{}); ok && len(v) > 0 {
		phoneNumbers, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.PhoneNumbers = *phoneNumbers
	}
	if v, ok := (*source)["photo_url"].(string); ok && len(v) > 0 {
		obj.PhotoUrl = &v
	}
	if v, ok := (*source)["recruiter"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateUsers(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Recruiter = &(*list)[0]
	}
	if v, ok := (*source)["social_media_addresses"].([]interface{}); ok && len(v) > 0 {
		addresses, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Addresses = *addresses
	}
	if v, ok := (*source)["tags"].([]interface{}); ok && len(v) > 0 {
		tags := *sliceItoSliceA(&v)
		obj.Tags = tags
	}
	if v, ok := (*source)["title"].(string); ok && len(v) > 0 {
		obj.Title = &v
	}
	if v, ok := (*source)["updated_at"].(string); ok && len(v) > 0 {
		obj.UpdatedAt = &v
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
	if v := item.Addresses; len(v) > 0 {
		converted := []greenhouse.TypeTypeValue(v)
		candidate["addresses"] = flattenTypeTypeValues(ctx, &converted)
	} else {
		candidate["addresses"] = emptyList()
	}
	candidate["application_ids"] = item.ApplicationIds
	candidate["attachments"] = flattenAttachments(ctx, &item.Attachments)
	if v := item.CanEmail; v != nil {
		candidate["can_email"] = *v
	}
	if v := item.Company; v != nil {
		candidate["company"] = *v
	}

	if v := item.Coordinator; v != nil {
		candidate["coordinator"] = []map[string]interface{}{flattenUser(ctx, v)}
	} else {
		candidate["coordinator"] = emptyList()
	}
	if v := item.CreatedAt; v != nil {
		candidate["created_at"] = *v
	}
	/* TODO
	if v := item.CustomFields; len(v) > 0 {

	}
	*/
	candidate["educations"] = flattenEducations(ctx, &item.Educations)
	if v := item.EmailAddresses; len(v) > 0 {
		converted := []greenhouse.TypeTypeValue(v)
		candidate["email_addresses"] = flattenTypeTypeValues(ctx, &converted)
	} else {
		candidate["email_addresses"] = emptyList()
	}
	candidate["employments"] = flattenEmployments(ctx, &item.Employments)
	if v := item.FirstName; v != nil {
		candidate["first_name"] = *v
	}
	if v := item.IsPrivate; v != nil {
		candidate["is_private"] = *v
	}
	/* TODO
	if v := item.KeyedCustomFields; len(v) > 0 {

	}
	*/
	if v := item.LastActivity; v != nil {
		candidate["last_activity"] = *v
	}
	if v := item.LastName; v != nil {
		candidate["last_name"] = *v
	}
	candidate["linked_user_ids"] = item.LinkedUserIds
	if v := item.PhoneNumbers; len(v) > 0 {
		converted := []greenhouse.TypeTypeValue(v)
		candidate["phone_numbers"] = flattenTypeTypeValues(ctx, &converted)
	} else {
		candidate["phone_numbers"] = emptyList()
	}
	if v := item.PhotoUrl; v != nil {
		candidate["photo_url"] = *v
	}
	if v := item.Recruiter; v != nil {
		candidate["recruiter"] = []map[string]interface{}{flattenUser(ctx, v)}
	} else {
		candidate["recruiter"] = emptyList()
	}
	if v := item.SocialMediaAddresses; v != nil {
		converted := []greenhouse.TypeTypeValue(v)
		candidate["social_media_addresses"] = flattenTypeTypeValues(ctx, &converted)
	} else {
		candidate["social_media_addresses"] = emptyList()
	}
	candidate["tags"] = item.Tags
	if v := item.Title; v != nil {
		candidate["title"] = *v
	}
	if v := item.UpdatedAt; v != nil {
		candidate["updated_at"] = *v
	}
	if v := item.WebsiteAddresses; len(v) > 0 {
		converted := []greenhouse.TypeTypeValue(v)
		candidate["website_addresses"] = flattenTypeTypeValues(ctx, &converted)
	} else {
		candidate["website_addresses"] = emptyList()
	}
	return candidate
}

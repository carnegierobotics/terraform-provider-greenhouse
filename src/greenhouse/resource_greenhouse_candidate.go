package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseCandidate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseCandidateCreate,
		ReadContext:   resourceGreenhouseCandidateRead,
		UpdateContext: resourceGreenhouseCandidateUpdate,
		DeleteContext: resourceGreenhouseCandidateDelete,
		Exists:        resourceGreenhouseCandidateExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseCandidate(),
	}
}

func resourceGreenhouseCandidateExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/candidates/%d", id))
}

func resourceGreenhouseCandidateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseCandidateCreate.")
	createObject := greenhouse.Candidate{
		FirstName:            d.Get("first_name").(string),
		LastName:             d.Get("last_name").(string),
		Company:              d.Get("company").(string),
		Title:                d.Get("title").(string),
		PhoneNumbers:         d.Get("phone_numbers").([]greenhouse.TypeTypeValue),
		Addresses:            d.Get("addresses").([]greenhouse.TypeTypeValue),
		EmailAddresses:       d.Get("email_addresses").([]greenhouse.TypeTypeValue),
		WebsiteAddresses:     d.Get("website_addresses").([]greenhouse.TypeTypeValue),
		SocialMediaAddresses: d.Get("social_media_addresses").([]greenhouse.TypeTypeValue),
		Educations:           d.Get("educations").([]greenhouse.Education),
		Employments:          d.Get("employments").([]greenhouse.Employment),
		Tags:                 d.Get("tags").([]string),
		CustomFields:         d.Get("custom_fields").(map[string]interface{}),
		ActivityFeedNotes:    d.Get("activity_feed_notes").([]greenhouse.ActivityFeed),
	}
	var err error
	var id int
	if d.Get("is_prospect").(bool) {
		createObject.Recruiter = d.Get("recruiter").(greenhouse.User)
		createObject.Coordinator = d.Get("coordinator").(greenhouse.User)
		createObject.Application = d.Get("application").(greenhouse.Application)
		id, err = greenhouse.CreateProspect(meta.(*greenhouse.Client), ctx, &createObject)
	} else {
		createObject.Applications = d.Get("applications").([]greenhouse.Application)
		id, err = greenhouse.CreateCandidate(meta.(*greenhouse.Client), ctx, &createObject)
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	tflog.Debug(ctx, "Kicking off resourceGreenhouseCandidateUpdate from resourceGreenhouseCandidateCreate.")
	return resourceGreenhouseCandidateUpdate(ctx, d, meta)
}

func resourceGreenhouseCandidateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseCandidateRead.")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetCandidate(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	convertedAddresses := []greenhouse.TypeTypeValue(obj.Addresses)
	d.Set("addresses", flattenTypeTypeValues(ctx, &convertedAddresses))
	d.Set("application_ids", obj.ApplicationIds)
	d.Set("attachments", flattenAttachments(ctx, &obj.Attachments))
	d.Set("company", obj.Company)
	d.Set("created_at", obj.CreatedAt)
	d.Set("first_name", obj.FirstName)
	d.Set("is_private", obj.IsPrivate)
	d.Set("last_activity", obj.LastActivity)
	d.Set("last_name", obj.LastName)
	convertedPhoneNumbers := []greenhouse.TypeTypeValue(obj.PhoneNumbers)
	d.Set("phone_numbers", flattenTypeTypeValues(ctx, &convertedPhoneNumbers))
	d.Set("photo_url", obj.PhotoUrl)
	d.Set("title", obj.Title)
	d.Set("updated_at", obj.UpdatedAt)
	tflog.Debug(ctx, "Finished resourceGreenhouseCandidateRead.")
	return nil
}

func resourceGreenhouseCandidateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseCandidateUpdate.")
	if d.HasChanges("educations") {

	}
	if d.HasChanges("employments") {

	}
	if d.HasChanges("attachments") {

	}
	if d.HasChanges("notes") {

	}
	if d.HasChanges("email_notes") {

	}
	tflog.Debug(ctx, "Kicking off resourceGreenhouseCandidateRead from resourceGreenhouseCandidateUpdate.")
	return resourceGreenhouseCandidateRead(ctx, d, meta)
}

func resourceGreenhouseCandidateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseCandidateDelete.")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteCandidate(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Debug(ctx, "Finished resourceGreenhouseCandidateDelete.")
	return nil
}

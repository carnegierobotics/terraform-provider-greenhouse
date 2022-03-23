package greenhouse

import (
	"context"
  "encoding/json"
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
  tflog.Debug(ctx, fmt.Sprintf("FirstName: %s", d.Get("first_name").(string)))
  createObj := greenhouse.Candidate{
	  FirstName: d.Get("first_name").(string),
    LastName: d.Get("last_name").(string),
    Company: d.Get("company").(string),
	  Title: d.Get("title").(string),
	  PhoneNumbers: *convertToTypeTypeValueList(d.Get("phone_numbers").(*schema.Set).List()),
	  Addresses: *convertToTypeTypeValueList(d.Get("addresses").(*schema.Set).List()),
	  EmailAddresses: *convertToTypeTypeValueList(d.Get("email_addresses").(*schema.Set).List()),
    WebsiteAddresses: *convertToTypeTypeValueList(d.Get("website_addresses").(*schema.Set).List()),
    SocialMediaAddresses: *convertToTypeTypeValueList(d.Get("social_media_addresses").(*schema.Set).List()),
	  Educations: *convertToEducationList(d.Get("educations").([]interface{})),
	  Employments: *convertToEmploymentList(d.Get("employments").([]interface{})),
	  Tags: *ConvertSliceInterfaceString(d.Get("tags").(*schema.Set).List()),
	  //CustomFields: d.Get("custom_fields").(*schema.Set)(map[string]interface{}),
	  ActivityFeedNotes: *convertToActivityFeedList(d.Get("activity_feed_notes").([]interface{})),
  }
  tflog.Debug(ctx, fmt.Sprintf("Initial candidate: %+v", createObj))
	var err error
  var diagErr diag.Diagnostics
	var id int
	if d.Get("is_prospect").(bool) {
    recruiter := d.Get("recruiter").([]interface{})
    if len(recruiter) == 1 {
      var recruiterObj greenhouse.User
		  diagErr = convertType(ctx, recruiter[0], &recruiterObj)
      if diagErr != nil {
        return diagErr
      }
      createObj.Recruiter = &recruiterObj
    }
    coordinator := d.Get("coordinator").([]interface{})
    if len(coordinator) == 1 {
      var coordinatorObj greenhouse.User
		  diagErr = convertType(ctx, coordinator[0], &coordinatorObj)
      if diagErr != nil {
        return diagErr
      }
      createObj.Coordinator = &coordinatorObj
    }
    application := d.Get("application").([]interface{})
    if len(application) == 1 {
      var applicationObj greenhouse.Application
		  diagErr = convertType(ctx, application[0], &applicationObj)
      if diagErr != nil {
        return diagErr
      }
      createObj.Application = &applicationObj
    }
    tflog.Debug(ctx, fmt.Sprintf("Creating prospect: %+v", createObj))
    jsonBody, err := json.Marshal(createObj)
    if err != nil {
      return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
    }
    tflog.Debug(ctx, fmt.Sprintf("JSON will be: %s", string(jsonBody)))
		id, err = greenhouse.CreateProspect(meta.(*greenhouse.Client), ctx, &createObj)
	} else {
		createObj.Applications = *convertToApplicationList(d.Get("applications").(*schema.Set).List())
    tflog.Debug(ctx, fmt.Sprintf("Creating candidate: %+v", createObj))
		id, err = greenhouse.CreateCandidate(meta.(*greenhouse.Client), ctx, &createObj)
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

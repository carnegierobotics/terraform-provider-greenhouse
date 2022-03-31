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
	var createObj greenhouse.Candidate
	if v, ok := d.Get("activity_feed_notes").([]interface{}); ok && len(v) < 0 {
		list, diagErr := inflateActivityFeeds(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.ActivityFeedNotes = *list
	}
	if v, ok := d.Get("addresses").([]interface{}); ok && len(v) > 0 {
		addresses, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.Addresses = *addresses
	}
	if v, ok := d.Get("company").(string); ok && len(v) > 0 {
		createObj.Company = v
	}
	if v, ok := d.Get("custom_fields").([]interface{}); ok && len(v) > 0 {
		createObj.CustomFields = v[0].(map[string]interface{})
	}
	if v, ok := d.Get("educations").([]interface{}); ok && len(v) > 0 {
		educations, diagErr := inflateEducations(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.Educations = *educations
	}
	if v, ok := d.Get("email_addresses").([]interface{}); ok && len(v) > 0 {
		emailAddresses, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.EmailAddresses = *emailAddresses
	}
	if v, ok := d.Get("employments").([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateEmployments(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.Employments = *list
	}
	if v, ok := d.Get("first_name").(string); ok && len(v) > 0 {
		createObj.FirstName = v
	}
	if v, ok := d.Get("last_name").(string); ok && len(v) > 0 {
		createObj.LastName = v
	}
	if v, ok := d.Get("phone_numbers").([]interface{}); ok && len(v) > 0 {
		phoneNumbers, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.PhoneNumbers = *phoneNumbers
	}
	if v, ok := d.Get("social_media_addresses").([]interface{}); ok && len(v) > 0 {
		socialMediaAddresses, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.SocialMediaAddresses = *socialMediaAddresses
	}
	if v, ok := d.Get("tags").([]interface{}); ok && len(v) > 0 {
		createObj.Tags = *sliceItoSliceA(&v)
	}
	if v, ok := d.Get("title").(string); ok && len(v) > 0 {
		createObj.Title = v
	}
	if v, ok := d.Get("website_addresses").([]interface{}); ok && len(v) > 0 {
		websiteAddresses, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.WebsiteAddresses = *websiteAddresses
	}
	tflog.Debug(ctx, fmt.Sprintf("Initial candidate: %+v", createObj))
	var err error
	var id int
	if d.Get("is_prospect").(bool) {
		if v, ok := d.Get("recruiter").([]interface{}); ok && len(v) == 1 {
			recruiterObj, diagErr := inflateUsers(ctx, &v)
			if diagErr != nil {
				return diagErr
			}
			if recruiterObj != nil && len(*recruiterObj) > 0 {
				createObj.Recruiter = &(*recruiterObj)[0]
			}
		}
		if v, ok := d.Get("coordinator").([]interface{}); ok && len(v) == 1 {
			coordinatorObj, diagErr := inflateUsers(ctx, &v)
			if diagErr != nil {
				return diagErr
			}
			if coordinatorObj != nil && len(*coordinatorObj) > 0 {
				createObj.Coordinator = &(*coordinatorObj)[0]
			}
		}
		if v, ok := d.Get("application").([]interface{}); ok && len(v) == 1 {
			applicationObj, diagErr := inflateApplications(ctx, &v)
			if diagErr != nil {
				tflog.Debug(ctx, "Error occurred during application inflation.")
				return diagErr
			}
			if applicationObj != nil && len(*applicationObj) > 0 {
				tflog.Debug(ctx, "Setting application.")
				app := (*applicationObj)[0]
				createObj.Application = &app
			}
		}
		tflog.Debug(ctx, fmt.Sprintf("Creating prospect: %+v", createObj))
		jsonBody, err := json.Marshal(createObj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		tflog.Debug(ctx, fmt.Sprintf("JSON will be: %s", string(jsonBody)))
		id, err = greenhouse.CreateProspect(meta.(*greenhouse.Client), ctx, &createObj)
	} else {
		if v, ok := d.Get("applications").([]interface{}); ok && len(v) > 0 {
			apps, diagErr := inflateApplications(ctx, &v)
			if diagErr != nil {
				return diagErr
			}
			createObj.Applications = *apps
		}
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
	for k, v := range flattenCandidate(ctx, obj) {
		d.Set(k, v)
	}
	tflog.Debug(ctx, "Finished resourceGreenhouseCandidateRead.")
	return nil
}

func resourceGreenhouseCandidateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseCandidateUpdate.")
	/*
			if d.HasChanges("educations") {
		    err := updateEducations(ctx, d, meta)
		    if err != nil {
		      return err
		    }
		  }
		 	if d.HasChanges("employments") {
		    o, n := d.GetChange("employments")
			}
			if d.HasChanges("attachments") {
		    o, n := d.GetChange("attachments")
			}
			if d.HasChanges("notes") {
		    o, n := d.GetChange("notes")
			}
			if d.HasChanges("email_notes") {
		    o, n := d.GetChange("email_notes")
			}
			if d.HasChanges("tags") {
		    o, n := d.GetChange("tags")
			}
	*/
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
	d.SetId("")
	return nil
}

func updateEducations(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var add *[]greenhouse.Education
	var del *[]greenhouse.Education
	var diagErr diag.Diagnostics
	o, n := d.GetChange("educations")
	v, ok1 := o.([]interface{})
	w, ok2 := n.([]interface{})
	if !ok1 || !ok2 {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Failed to convert to []interface{}"}}
	}
	if len(v) == 0 {
		add, diagErr = inflateEducations(ctx, &w)
		if diagErr != nil {
			return diagErr
		}
	} else if len(w) == 0 {
		del, diagErr = inflateEducations(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
	} else {
		inflatedv, diagErr := inflateEducations(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		inflatedw, diagErr := inflateEducations(ctx, &w)
		if diagErr != nil {
			return diagErr
		}
		add, del = findAddDelete(inflatedv, inflatedw)
	}
	for _, edu := range *add {
		err = greenhouse.AddEducationToCandidate(meta.(*greenhouse.Client), ctx, cId, &edu)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	for _, edu := range *del {
		err = greenhouse.DeleteEducationFromCandidate(meta.(*greenhouse.Client), ctx, cId, edu.Id)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return nil
}

func findAddDelete(v *[]greenhouse.Education, w *[]greenhouse.Education) (*[]greenhouse.Education, *[]greenhouse.Education) {
	add := make([]greenhouse.Education, 0)
	del := make([]greenhouse.Education, 0)
	for _, i1 := range *v {
		match := false
		for _, i2 := range *w {
			if i1.Id == i2.Id {
				match = true
				break
			}
		}
		if !match {
			del = append(del, i1)
		}
	}
	for _, i1 := range *w {
		match := false
		for _, i2 := range *v {
			if i1.Id == i2.Id {
				match = true
				break
			}
		}
		if !match {
			add = append(add, i1)
		}
	}
	return &add, &del
}

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
			StateContext: resourceGreenhouseCandidateImport,
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
	tflog.Trace(ctx, "Started resourceGreenhouseCandidateCreate.")
	tflog.Trace(ctx, fmt.Sprintf("FirstName: %s", d.Get("first_name").(string)))
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
		createObj.Company = &v
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
		createObj.FirstName = &v
	}
	if v, ok := d.Get("last_name").(string); ok && len(v) > 0 {
		createObj.LastName = &v
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
		createObj.Title = &v
	}
	if v, ok := d.Get("website_addresses").([]interface{}); ok && len(v) > 0 {
		websiteAddresses, diagErr := inflateTypeTypeValues(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
		createObj.WebsiteAddresses = *websiteAddresses
	}
	tflog.Trace(ctx, fmt.Sprintf("Initial candidate: %+v", createObj))
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
				tflog.Trace(ctx, "Error occurred during application inflation.")
				return diagErr
			}
			if applicationObj != nil && len(*applicationObj) > 0 {
				tflog.Trace(ctx, "Setting application.")
				app := (*applicationObj)[0]
				createObj.Application = &app
			}
		}
		tflog.Trace(ctx, fmt.Sprintf("Creating prospect: %+v", createObj))
		diagErr := logJson(ctx, "resourceGreenhouseCandidateCreate", createObj)
		if diagErr != nil {
			return diagErr
		}
		id, err = greenhouse.CreateProspect(meta.(*greenhouse.Client), ctx, &createObj)
	} else {
		if v, ok := d.Get("applications").([]interface{}); ok && len(v) > 0 {
			apps, diagErr := inflateApplications(ctx, &v)
			if diagErr != nil {
				return diagErr
			}
			createObj.Applications = *apps
		}
		tflog.Trace(ctx, fmt.Sprintf("Creating candidate: %+v", createObj))
		id, err = greenhouse.CreateCandidate(meta.(*greenhouse.Client), ctx, &createObj)
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	tflog.Trace(ctx, "Kicking off resourceGreenhouseCandidateUpdate from resourceGreenhouseCandidateCreate.")
	return resourceGreenhouseCandidateUpdate(ctx, d, meta)
}

func resourceGreenhouseCandidateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Started resourceGreenhouseCandidateRead.")
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
	tflog.Trace(ctx, "Finished resourceGreenhouseCandidateRead.")
	return nil
}

func resourceGreenhouseCandidateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Started resourceGreenhouseCandidateUpdate.")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if d.HasChanges("educations") {
		err := updateEducations(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	if d.HasChanges("employments") {
		err := updateEmployments(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	if d.HasChanges("attachments") {
		err := updateAttachments(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	if d.HasChanges("notes") {
		err := updateNotes(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	if d.HasChanges("email_notes") {
		err := updateEmailNotes(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	if d.HasChanges("tags") {
		err := updateTags(ctx, d, meta)
		if err != nil {
			return err
		}
	}
	if d.HasChanges("anonymize") {
		if v, ok := d.Get("anonymize").([]string); ok && len(v) > 0 {
			_, err := greenhouse.AnonymizeCandidate(meta.(*greenhouse.Client), ctx, id, v)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		}
	}
	if d.HasChanges("merge") {
		v, ok := d.Get("merge").(int)
		if !ok {
			return diag.Diagnostics{{Severity: diag.Error, Summary: "Could not get ID to merge to."}}
		}
		err := greenhouse.MergeCandidates(meta.(*greenhouse.Client), ctx, v, id)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		d.SetId("")
		return nil
	}
	tflog.Trace(ctx, "Kicking off resourceGreenhouseCandidateRead from resourceGreenhouseCandidateUpdate.")
	return resourceGreenhouseCandidateRead(ctx, d, meta)
}

func resourceGreenhouseCandidateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Started resourceGreenhouseCandidateDelete.")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteCandidate(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Trace(ctx, "Finished resourceGreenhouseCandidateDelete.")
	d.SetId("")
	return nil
}

func resourceGreenhouseCandidateImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseCandidateRead)
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
	addI, delI, diagErr := findAddDelete(ctx, o, n)
	if diagErr != nil {
		return diagErr
	}
	add, diagErr = inflateEducations(ctx, addI)
	if diagErr != nil {
		return diagErr
	}
	del, diagErr = inflateEducations(ctx, delI)
	if diagErr != nil {
		return diagErr
	}
	for _, obj := range *add {
		err = greenhouse.AddEducationToCandidate(meta.(*greenhouse.Client), ctx, cId, &obj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	for _, obj := range *del {
		if v := obj.Id; v != nil {
			err = greenhouse.DeleteEducationFromCandidate(meta.(*greenhouse.Client), ctx, cId, *v)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		}
	}
	return nil
}

func updateEmployments(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var add *[]greenhouse.Employment
	var del *[]greenhouse.Employment
	var diagErr diag.Diagnostics
	o, n := d.GetChange("employments")
	addI, delI, diagErr := findAddDelete(ctx, o, n)
	if diagErr != nil {
		return diagErr
	}
	add, diagErr = inflateEmployments(ctx, addI)
	if diagErr != nil {
		return diagErr
	}
	del, diagErr = inflateEmployments(ctx, delI)
	if diagErr != nil {
		return diagErr
	}
	for _, obj := range *add {
		err = greenhouse.AddEmploymentToCandidate(meta.(*greenhouse.Client), ctx, cId, &obj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	for _, obj := range *del {
		if v := obj.Id; v != nil {
			err = greenhouse.DeleteEmploymentFromCandidate(meta.(*greenhouse.Client), ctx, cId, *v)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		}
	}
	return nil
}

func updateAttachments(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var add *[]greenhouse.Attachment
	var diagErr diag.Diagnostics
	o, n := d.GetChange("attachments")
	addI, _, diagErr := findAddDelete(ctx, o, n)
	if diagErr != nil {
		return diagErr
	}
	add, diagErr = inflateAttachments(ctx, addI)
	if diagErr != nil {
		return diagErr
	}
	for _, obj := range *add {
		err = greenhouse.AddAttachmentToCandidate(meta.(*greenhouse.Client), ctx, cId, &obj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return nil
}

func updateNotes(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var add *[]greenhouse.Note
	var diagErr diag.Diagnostics
	o, n := d.GetChange("notes")
	addI, _, diagErr := findAddDelete(ctx, o, n)
	if diagErr != nil {
		return diagErr
	}
	add, diagErr = inflateNotes(ctx, addI)
	if diagErr != nil {
		return diagErr
	}
	for _, obj := range *add {
		err = greenhouse.AddNoteToCandidate(meta.(*greenhouse.Client), ctx, cId, &obj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return nil
}

func updateEmailNotes(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var add *[]greenhouse.Email
	var diagErr diag.Diagnostics
	o, n := d.GetChange("email_notes")
	addI, _, diagErr := findAddDelete(ctx, o, n)
	if diagErr != nil {
		return diagErr
	}
	add, diagErr = inflateEmails(ctx, addI)
	if diagErr != nil {
		return diagErr
	}
	for _, obj := range *add {
		err = greenhouse.AddEmailNoteToCandidate(meta.(*greenhouse.Client), ctx, cId, &obj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return nil
}

func updateTags(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	cId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var add []string
	var del []string
	oi, ni := d.GetChange("tags")
	o, ok1 := oi.([]interface{})
	n, ok2 := ni.([]interface{})
	if !ok1 || !ok2 {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Failed to convert to []interface{}"}}
	}
	for _, v := range o {
		match := false
		for _, w := range n {
			if v.(string) == w.(string) {
				match = true
				break
			}
		}
		if !match {
			del = append(del, v.(string))
		}
	}
	for _, v := range n {
		match := false
		for _, w := range o {
			if v.(string) == w.(string) {
				match = true
				break
			}
		}
		if !match {
			add = append(add, v.(string))
		}
	}
	allTags, err := greenhouse.GetAllCandidateTags(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for _, name := range add {
		for _, tag := range *allTags {
			if *tag.Name == name {
				err = greenhouse.CreateTagForCandidate(meta.(*greenhouse.Client), ctx, cId, *tag.Id)
				if err != nil {
					return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
				}
			}
		}
	}
	for _, name := range del {
		for _, tag := range *allTags {
			if *tag.Name == name {
				err = greenhouse.DeleteTagFromCandidate(meta.(*greenhouse.Client), ctx, cId, *tag.Id)
				if err != nil {
					return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
				}
			}
		}
	}
	return nil
}

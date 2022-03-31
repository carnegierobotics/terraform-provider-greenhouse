package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"reflect"
	"strconv"
)

func resourceGreenhouseJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseJobCreate,
		ReadContext:   resourceGreenhouseJobRead,
		UpdateContext: resourceGreenhouseJobUpdate,
		DeleteContext: resourceGreenhouseJobDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: schemaGreenhouseJob(),
	}
}

func resourceGreenhouseJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseJobCreate")
	createObject := greenhouse.JobCreateInfo{
		TemplateJobId:  d.Get("template_job_id").(int),
		NumberOpenings: d.Get("number_of_openings").(int),
		JobPostName:    d.Get("job_post_name").(string),
		JobName:        d.Get("job_name").(string),
		DepartmentId:   d.Get("department_id").(int),
		RequisitionId:  d.Get("requisition_id").(string),
	}
	if v, ok := d.Get("office_ids").([]interface{}); ok && len(v) > 0 {
		createObject.OfficeIds = *sliceItoSliceD(&v)
	}
	if v, ok := d.Get("opening_ids").([]interface{}); ok && len(v) > 0 {
		createObject.OpeningIds = *sliceItoSliceA(&v)
	}
	id, err := greenhouse.CreateJob(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	tflog.Debug(ctx, "Kicking off resourceGreenhouseJobUpdate from resourceGreenhouseJobCreate")
	return resourceGreenhouseJobUpdate(ctx, d, meta)
}

func resourceGreenhouseJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseJobRead")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetJob(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Debug(ctx, "Debugging job", "job", fmt.Sprintf("%+v", obj))
	for k, v := range flattenJob(ctx, obj) {
		d.Set(k, v)
	}
	tflog.Debug(ctx, "Finished resourceGreenhouseJobRead")
	return nil
}

func resourceGreenhouseJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseJobUpdate")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if d.HasChange("openings") {
		diagErr := updateOpenings(ctx, d, meta)
		if diagErr != nil {
			return diagErr
		}
	}
	if d.HasChange("hiring_team") {
		teamUpdateObject, err := transformHiringTeam(ctx, d.Get("hiring_team"))
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		err = greenhouse.UpdateJobHiringTeam(meta.(*greenhouse.Client), ctx, id, &teamUpdateObject)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	updateObject := greenhouse.JobUpdateInfo{
		Name:                    d.Get("job_name").(string),
		Notes:                   d.Get("notes").(string),
		Anywhere:                d.Get("anywhere").(bool),
		RequisitionId:           d.Get("requisition_id").(string),
		TeamandResponsibilities: d.Get("team_and_responsibilities").(string),
		HowToSellThisJob:        d.Get("how_to_sell_this_job").(string),
		DepartmentId:            d.Get("department_id").(int),
	}
	if v, ok := d.Get("office_ids").([]interface{}); ok && len(v) > 0 {
		updateObject.OfficeIds = *sliceItoSliceD(&v)
	}
	err = greenhouse.UpdateJob(meta.(*greenhouse.Client), ctx, id, &updateObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Debug(ctx, "Kicking off resourceGreenhouseJobRead from resourceGreenhouseJobUpdate")
	return resourceGreenhouseJobRead(ctx, d, meta)
}

func transformHiringTeam(ctx context.Context, hiringTeam interface{}) (map[string][]greenhouse.HiringMemberUpdateInfo, error) {
	update := make(map[string][]greenhouse.HiringMemberUpdateInfo)
	for _, team := range hiringTeam.([]interface{}) {
		teamItem := team.(map[string]interface{})
		teamName := teamItem["name"].(string)
		members := teamItem["members"].([]interface{})
		update[teamName] = make([]greenhouse.HiringMemberUpdateInfo, len(members), len(members))
		for j, member := range members {
			var obj greenhouse.HiringMemberUpdateInfo
			marshaled, err := json.Marshal(member)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(marshaled, &obj)
			if err != nil {
				return nil, err
			}
			update[teamName][j] = obj
		}
	}
	tflog.Debug(ctx, "Updating hiring team", "updateObj", fmt.Sprintf("%+v", update))
	return update, nil
}

func resourceGreenhouseJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for jobs."}}
}

func updateOpenings(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	jobId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var oldList *[]greenhouse.JobOpening
	var newList *[]greenhouse.JobOpening
	var add *greenhouse.JobOpeningCreateInfo
	var del *[]int
	var upd *[]greenhouse.JobOpeningUpdateInfo
	var diagErr diag.Diagnostics
	o, n := d.GetChange("openings")
	if v, ok1 := o.([]interface{}); ok1 && len(v) > 0 {
		oldList, diagErr = inflateJobOpenings(ctx, &v)
		if diagErr != nil {
			return diagErr
		}
	}
	if w, ok2 := n.([]interface{}); ok2 && len(w) > 0 {
		newList, diagErr = inflateJobOpenings(ctx, &w)
		if diagErr != nil {
			return diagErr
		}
	}
	for _, i := range *oldList {
		obj, match := compareJobOpenings(ctx, &i, newList)
		if obj == nil && match {
			continue
		} else if match {
			updateObj := greenhouse.JobOpeningUpdateInfo{
				CloseReasonId: obj.CloseReason.Id,
				CustomFields:  []map[string]string{(*obj).CustomFields},
				Status:        obj.Status,
			}
			*upd = append(*upd, updateObj)
		} else {
			*del = append(*del, i.Id)
		}
	}
	for _, i := range *newList {
		_, match := compareJobOpenings(ctx, &i, oldList)
		if !match {
			createObj := greenhouse.Opening{
				CustomFields: []map[string]string{i.CustomFields},
			}
			(*add).Openings = append((*add).Openings, createObj)
		}
	}
	if add != nil {
		_, err := greenhouse.CreateJobOpenings(meta.(*greenhouse.Client), ctx, jobId, *add)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	if len(*upd) > 0 {
		for _, item := range *upd {
			err := greenhouse.UpdateJobOpenings(meta.(*greenhouse.Client), ctx, jobId, item.Id, &item)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		}
	}
	if len(*del) > 0 {
		err := greenhouse.DeleteJobOpenings(meta.(*greenhouse.Client), ctx, jobId, *del)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return nil
}

func compareJobOpenings(ctx context.Context, o *greenhouse.JobOpening, j *[]greenhouse.JobOpening) (*greenhouse.JobOpening, bool) {
	for _, item := range *j {
		if (*o).Id == item.Id {
			if reflect.DeepEqual(o, item) {
				return nil, true
			}
			return &item, true
		}
	}
	return nil, false
}

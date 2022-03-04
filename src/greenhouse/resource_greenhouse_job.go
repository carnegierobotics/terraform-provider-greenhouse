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

func resourceGreenhouseJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseJobCreate,
		ReadContext:   resourceGreenhouseJobRead,
		UpdateContext: resourceGreenhouseJobUpdate,
		DeleteContext: resourceGreenhouseJobDelete,
		Exists: resourceGreenhouseJobExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseJob(),
	}
}

func resourceGreenhouseJobExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "jobs", id, context.TODO())
}

func resourceGreenhouseJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.JobCreateInfo{
		TemplateJobId:  d.Get("template_job_id").(int),
		NumberOpenings: d.Get("number_of_openings").(int),
		JobPostName:    d.Get("job_post_name").(string),
		JobName:        d.Get("job_name").(string),
		DepartmentId:   d.Get("department_id").(int),
		OfficeIds:      convertListIToListD(d.Get("office_ids").(*schema.Set).List()),
		RequisitionId:  d.Get("requisition_id").(string),
		OpeningIds:     convertListIToListA(d.Get("opening_ids").(*schema.Set).List()),
	}
	id, err := greenhouse.CreateJob(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseJobRead(ctx, d, meta)
}

func convertListIToListD(list []interface{}) []int {
	newList := make([]int, len(list))
	if len(list) == 0 {
		return newList
	}
	for i := range list {
		newList[i] = list[i].(int)
	}
	return newList
}

func convertListIToListA(list []interface{}) []string {
	newList := make([]string, len(list))
	if len(list) == 0 {
		return newList
	}
	for i := range list {
		newList[i] = list[i].(string)
	}
	return newList
}

func resourceGreenhouseJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	obj, err := greenhouse.GetJob(meta.(*greenhouse.Client), id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
  tflog.Debug(ctx, "Debugging job", "job", fmt.Sprintf("%+v", obj))
	d.Set("job_name", obj.Name)
	d.Set("departments", flattenDepartments(&obj.Departments))
	d.Set("offices", obj.Offices)
	d.Set("requisition_id", obj.RequisitionId)
	d.Set("openings", obj.Openings)
	d.Set("hiring_team", obj.HiringTeam)
	d.Set("notes", obj.Notes)
	d.Set("confidential", obj.Confidential)
	d.Set("status", obj.Status)
	d.Set("created_at", obj.CreatedAt)
	d.Set("opened_at", obj.OpenedAt)
	d.Set("closed_at", obj.ClosedAt)
	d.Set("updated_at", obj.UpdatedAt)
	d.Set("is_template", obj.IsTemplate)
	d.Set("copied_from_id", obj.CopiedFromId)
	return nil
}

func resourceGreenhouseJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	if d.HasChange("hiring_team") {
		teamUpdateObject := convertHiringTeam(d.Get("hiring_team").(map[string][]interface{}))
		err = greenhouse.UpdateJobHiringTeam(meta.(*greenhouse.Client), id, &teamUpdateObject, context.TODO())
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
		}
	}
	updateObject := greenhouse.JobUpdateInfo{
		Name:                     d.Get("job_name").(string),
		Notes:                    d.Get("notes").(string),
		Anywhere:                 d.Get("anywhere").(bool),
		RequisitionId:            d.Get("requisition_id").(string),
		TeamsandResponsibilities: d.Get("teams_and_responsibilities").(string),
		HowToSellThisJob:         d.Get("how_to_sell_this_job").(string),
		OfficeIds:                convertListIToListD(d.Get("office_ids").(*schema.Set).List()),
		DepartmentId:             d.Get("department_id").(int),
	}
	err = greenhouse.UpdateJob(meta.(*greenhouse.Client), id, &updateObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	return resourceGreenhouseJobRead(ctx, d, meta)
}

func convertHiringTeam(list map[string][]interface{}) map[string][]greenhouse.HiringMemberUpdateInfo {
	var newMap map[string][]greenhouse.HiringMemberUpdateInfo
	for k, v := range list {
		newMap[k] = make([]greenhouse.HiringMemberUpdateInfo, len(v))
		for i := range v {
			newMap[k][i] = list[k][i].(greenhouse.HiringMemberUpdateInfo)
		}
	}
	return newMap
}

func resourceGreenhouseJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for jobs."},}
}

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
		OfficeIds:      convertListIToListD(d.Get("office_ids").(*schema.Set).List()),
		RequisitionId:  d.Get("requisition_id").(string),
		OpeningIds:     convertListIToListA(d.Get("opening_ids").(*schema.Set).List()),
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
	d.Set("job_name", obj.Name)
	d.Set("departments", flattenDepartments(&obj.Departments))
	d.Set("offices", flattenOffices(&obj.Offices))
	d.Set("requisition_id", obj.RequisitionId)
	d.Set("openings", flattenJobOpenings(&obj.Openings))
	//d.Set("hiring_team", flattenHiringTeam(ctx, &obj.HiringTeam))
	tflog.Debug(ctx, "Hiring team after flattening", "team", fmt.Sprintf("%+v", d.Get("hiring_team")))
	d.Set("notes", obj.Notes)
	d.Set("confidential", obj.Confidential)
	d.Set("status", obj.Status)
	d.Set("created_at", obj.CreatedAt)
	d.Set("opened_at", obj.OpenedAt)
	d.Set("closed_at", obj.ClosedAt)
	d.Set("updated_at", obj.UpdatedAt)
	d.Set("is_template", obj.IsTemplate)
	d.Set("copied_from_id", obj.CopiedFromId)
	d.Set("custom_fields", obj.CustomFields)
	//d.Set("keyed_custom_fields", obj.KeyedCustomFields)
	tflog.Debug(ctx, "Finished resourceGreenhouseJobRead")
	return nil
}

func resourceGreenhouseJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Started resourceGreenhouseJobUpdate")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if d.HasChanges("hiring_team") {
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
		OfficeIds:               convertListIToListD(d.Get("office_ids").(*schema.Set).List()),
		DepartmentId:            d.Get("department_id").(int),
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
	for _, team := range hiringTeam.(*schema.Set).List() {
		teamItem := team.(map[string]interface{})
		teamName := teamItem["name"].(string)
		members := teamItem["members"].(*schema.Set).List()
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

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
			StateContext: resourceGreenhouseJobImport,
		},
		Schema: schemaGreenhouseJob(),
	}
}

func resourceGreenhouseJobCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Started resourceGreenhouseJobCreate")
	var obj greenhouse.JobCreateInfo
	if v, ok := d.Get("template_job_id").(int); ok {
		obj.TemplateJobId = &v
	}
	if v, ok := d.Get("number_of_openings").(int); ok {
		obj.NumberOpenings = &v
	}
	if v, ok := d.Get("job_post_name").(string); ok && len(v) > 0 {
		obj.JobPostName = &v
	}
	if v, ok := d.Get("job_name").(string); ok && len(v) > 0 {
		obj.JobName = &v
	}
	if v, ok := d.Get("department_id").(int); ok {
		obj.DepartmentId = &v
	}
	if v, ok := d.Get("requisition_id").(string); ok {
		obj.RequisitionId = &v
	}
	if v, ok := d.Get("office_ids").([]interface{}); ok && len(v) > 0 {
		obj.OfficeIds = *sliceItoSliceD(&v)
	}
	if v, ok := d.Get("opening_ids").([]interface{}); ok && len(v) > 0 {
		obj.OpeningIds = *sliceItoSliceA(&v)
	}
	id, err := greenhouse.CreateJob(meta.(*greenhouse.Client), ctx, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	tflog.Trace(ctx, "Kicking off resourceGreenhouseJobUpdate from resourceGreenhouseJobCreate")
	return resourceGreenhouseJobUpdate(ctx, d, meta)
}

func resourceGreenhouseJobRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Started resourceGreenhouseJobRead")
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetJob(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Trace(ctx, "Debugging job", "job", fmt.Sprintf("%+v", obj))
	for k, v := range flattenJob(ctx, obj) {
		d.Set(k, v)
	}
	tflog.Trace(ctx, "Finished resourceGreenhouseJobRead")
	return nil
}

func resourceGreenhouseJobUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Started resourceGreenhouseJobUpdate")
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
	var obj greenhouse.JobUpdateInfo
	if v, ok := d.Get("name").(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	if v, ok := d.Get("notes").(string); ok && len(v) > 0 {
		obj.Notes = &v
	}
	if v, ok := d.Get("anywhere").(bool); ok {
		obj.Anywhere = &v
	}
	if v, ok := d.Get("requisition_id").(string); ok && len(v) > 0 {
		obj.RequisitionId = &v
	}
	if v, ok := d.Get("team_and_responsibilities").(string); ok && len(v) > 0 {
		obj.TeamandResponsibilities = &v
	}
	if v, ok := d.Get("how_to_sell_this_job").(string); ok && len(v) > 0 {
		obj.HowToSellThisJob = &v
	}
	if v, ok := d.Get("department_id").(int); ok {
		obj.DepartmentId = &v
	}
	if v, ok := d.Get("office_ids").([]interface{}); ok && len(v) > 0 {
		obj.OfficeIds = *sliceItoSliceD(&v)
	}
	err = greenhouse.UpdateJob(meta.(*greenhouse.Client), ctx, id, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Trace(ctx, "Kicking off resourceGreenhouseJobRead from resourceGreenhouseJobUpdate")
	return resourceGreenhouseJobRead(ctx, d, meta)
}

func resourceGreenhouseJobDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for jobs."}}
}

func resourceGreenhouseJobImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseJobRead)
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
			*del = append(*del, Int(i.Id))
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
			err := greenhouse.UpdateJobOpenings(meta.(*greenhouse.Client), ctx, jobId, Int(item.Id), &item)
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

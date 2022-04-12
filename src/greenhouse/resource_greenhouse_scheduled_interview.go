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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseScheduledInterview() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseScheduledInterviewCreate,
		ReadContext:   resourceGreenhouseScheduledInterviewRead,
		UpdateContext: resourceGreenhouseScheduledInterviewUpdate,
		DeleteContext: resourceGreenhouseScheduledInterviewDelete,
		Exists:        resourceGreenhouseScheduledInterviewExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseScheduledInterviewImport,
		},
		Schema: schemaGreenhouseScheduledInterview(),
	}
}

func resourceGreenhouseScheduledInterviewExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/scheduled_interviews/%d", id))
}

func resourceGreenhouseScheduledInterviewCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var obj greenhouse.ScheduledInterviewCreateInfo
	if v, ok := d.Get("application_id").(int); ok {
		obj.ApplicationId = &v
	}
	if v, ok := d.Get("end").([]interface{}); ok && len(v) > 0 {
		list, err := inflateScheduledInterviewDates(ctx, &v)
		if err != nil {
			return err
		}
		obj.End = (*list)[0].DateTime
	}
	if v, ok := d.Get("external_event_id").(string); ok && len(v) > 0 {
		obj.ExternalEventId = &v
	}
	if v, ok := d.Get("interview_id").(int); ok {
		obj.InterviewId = &v
	}
	if v, ok := d.Get("interviewers").([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewers(ctx, &v)
		if err != nil {
			return err
		}
		obj.Interviewers = *list
	}
	if v, ok := d.Get("start").([]interface{}); ok && len(v) > 0 {
		list, err := inflateScheduledInterviewDates(ctx, &v)
		if err != nil {
			return err
		}
		obj.Start = (*list)[0].DateTime
	}
	id, err := greenhouse.CreateScheduledInterview(meta.(*greenhouse.Client), ctx, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	return resourceGreenhouseScheduledInterviewUpdate(ctx, d, meta)
}

func resourceGreenhouseScheduledInterviewRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetScheduledInterview(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenScheduledInterview(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseScheduledInterviewUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	var obj greenhouse.ScheduledInterviewUpdateInfo
	if d.HasChanges("end") {
		if v, ok := d.Get("end").([]interface{}); ok && len(v) > 0 {
			list, err := inflateScheduledInterviewDates(ctx, &v)
			if err != nil {
				return err
			}
			obj.End = (*list)[0].DateTime
		}
	}
	if d.HasChanges("external_event_id") {
		if v, ok := d.Get("external_event_id").(string); ok && len(v) > 0 {
			obj.ExternalEventId = &v
		}
	}
	if d.HasChanges("interviewers") {
		if v, ok := d.Get("interviewers").([]interface{}); ok && len(v) > 0 {
			list, err := inflateInterviewers(ctx, &v)
			if err != nil {
				return err
			}
			obj.Interviewers = *list
		}
	}
	if d.HasChanges("location") {
		if v, ok := d.Get("location").(string); ok && len(v) > 0 {
			obj.Location = &v
		}
	}
	if d.HasChanges("start") {
		if v, ok := d.Get("start").([]interface{}); ok && len(v) > 0 {
			list, err := inflateScheduledInterviewDates(ctx, &v)
			if err != nil {
				return err
			}
			obj.Start = (*list)[0].DateTime
		}
	}
	err = greenhouse.UpdateScheduledInterview(meta.(*greenhouse.Client), ctx, id, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return resourceGreenhouseCandidateRead(ctx, d, meta)
}

func resourceGreenhouseScheduledInterviewDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteScheduledInterview(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("")
	return nil
}

func resourceGreenhouseScheduledInterviewImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseScheduledInterviewRead)
}

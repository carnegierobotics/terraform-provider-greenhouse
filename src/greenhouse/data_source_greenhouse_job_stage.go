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
	"strconv"
)

func dataSourceGreenhouseJobStage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseJobStageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseJobStageRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, ok := d.GetOk("id")
	var stage *greenhouse.JobStage
	var err error
	if ok {
		stage, err = greenhouse.GetJobStage(meta.(*greenhouse.Client), ctx, id.(int))
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	} else {
		job_id := d.Get("job_id").(int)
		name := d.Get("name").(string)
		list, err := greenhouse.GetJobStagesForJob(meta.(*greenhouse.Client), ctx, job_id)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		for _, item := range *list {
			if name == *item.Name {
				stage = &item
			}
		}
	}
	d.SetId(strconv.Itoa(*stage.Id))
	d.Set("created_at", stage.CreatedAt)
	d.Set("interviews", flattenInterviews(ctx, &stage.Interviews))
	d.Set("job_id", stage.JobId)
	d.Set("name", stage.Name)
	d.Set("priority", stage.Priority)
	d.Set("updated_at", stage.UpdatedAt)
	return nil
}

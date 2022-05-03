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

func resourceGreenhouseJobOpening() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseJobOpeningCreate,
		ReadContext:   resourceGreenhouseJobOpeningRead,
		UpdateContext: resourceGreenhouseJobOpeningUpdate,
		DeleteContext: resourceGreenhouseJobOpeningDelete,
		Exists:        resourceGreenhouseJobOpeningExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseJobOpeningImport,
		},
		Schema: schemaGreenhouseJobOpening(),
	}
}

func resourceGreenhouseJobOpeningExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/job_openings/%d", id))
}

func resourceGreenhouseJobOpeningCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Create is not supported for job_openings."}}
}

func resourceGreenhouseJobOpeningRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if v, ok := d.Get("job_id").(int); ok {
		obj, err := greenhouse.GetJobOpening(meta.(*greenhouse.Client), ctx, v, id)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		d.Set("application_id", obj.ApplicationId)
		d.Set("close_reason", flattenCloseReason(ctx, obj.CloseReason))
		return nil
	}
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting job_opening."}}
}

func resourceGreenhouseJobOpeningUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for job_openings."}}
}

func resourceGreenhouseJobOpeningDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for job_openings."}}
}

func resourceGreenhouseJobOpeningImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseJobOpeningRead)
}

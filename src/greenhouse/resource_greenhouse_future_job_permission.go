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

func resourceGreenhouseFutureJobPermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseFutureJobPermissionCreate,
		ReadContext:   resourceGreenhouseFutureJobPermissionRead,
		UpdateContext: resourceGreenhouseFutureJobPermissionUpdate,
		DeleteContext: resourceGreenhouseFutureJobPermissionDelete,
		Exists:        resourceGreenhouseFutureJobPermissionExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseFutureJobPermissionImport,
		},
		Schema: schemaGreenhouseFutureJobPermission(),
	}
}

func resourceGreenhouseFutureJobPermissionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/xxx/%d", id))
}

func resourceGreenhouseFutureJobPermissionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var obj greenhouse.FutureJobPermission
	if v, ok := d.Get("department_id").(int); ok {
		obj.DepartmentId = &v
	}
	if v, ok := d.Get("external_department_id").(string); ok && len(v) > 0 {
		obj.ExternalDepartmentId = &v
	}
	if v, ok := d.Get("external_office_id").(string); ok {
		obj.ExternalOfficeId = &v
	}
	if v, ok := d.Get("office_id").(int); ok {
		obj.OfficeId = &v
	}
	if v, ok := d.Get("user_role_id").(int); ok {
		obj.UserRoleId = &v
	}
	if v, ok := d.Get("user_id").(int); ok {
		id, err := greenhouse.CreateFutureJobPermission(meta.(*greenhouse.Client), ctx, v, &obj)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
		d.SetId(strconv.Itoa(id))
		return resourceGreenhouseFutureJobPermissionRead(ctx, d, meta)
	}
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Could not create permission."}}
}

func resourceGreenhouseFutureJobPermissionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	userId, ok := d.Get("user_id").(int)
	if !ok {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting user_id."}}
	}
	obj, err := greenhouse.GetFutureJobPermission(meta.(*greenhouse.Client), ctx, userId, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.Set("department_id", obj.DepartmentId)
	d.Set("external_department_id", obj.ExternalDepartmentId)
	d.Set("external_office_id", obj.ExternalOfficeId)
	d.Set("office_id", obj.OfficeId)
	d.Set("user_role_id", obj.UserRoleId)
	return nil
}

func resourceGreenhouseFutureJobPermissionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for xxx."}}
}

func resourceGreenhouseFutureJobPermissionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	permId, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	jobId, ok := d.Get("job_id").(int)
	if !ok {
		return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting job_id."}}
	}
	err = greenhouse.DeleteFutureJobPermission(meta.(*greenhouse.Client), ctx, jobId, permId)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("")
	return nil
}

func resourceGreenhouseFutureJobPermissionImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseFutureJobPermissionRead)
}

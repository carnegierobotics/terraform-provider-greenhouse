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

type ReadFunc func(d *schema.ResourceData, m interface{}) error

func resourceGreenhouseDepartment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseDepartmentCreate,
		ReadContext:   resourceGreenhouseDepartmentRead,
		UpdateContext: resourceGreenhouseDepartmentUpdate,
		DeleteContext: resourceGreenhouseDepartmentDelete,
		Exists:        resourceGreenhouseDepartmentExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseDepartmentImport,
		},
		Schema: schemaGreenhouseDepartment(),
	}
}

func resourceGreenhouseDepartmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/departments/%d", id))
}

func resourceGreenhouseDepartmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.Department{
		ExternalId:                 StringPtr(d.Get("external_id").(string)),
		Name:                       StringPtr(d.Get("name").(string)),
		ParentDepartmentExternalId: StringPtr(d.Get("external_parent_id").(string)),
		ParentId:                   IntPtr(d.Get("parent_id").(int)),
	}
	id, err := greenhouse.CreateDepartment(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseDepartmentUpdate(ctx, d, meta)
}

func resourceGreenhouseDepartmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetDepartment(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenDepartment(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseDepartmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	updateObject := greenhouse.Department{
		ExternalId: StringPtr(d.Get("external_id").(string)),
		Name:       StringPtr(d.Get("name").(string)),
	}
	err = greenhouse.UpdateDepartment(meta.(*greenhouse.Client), ctx, id, &updateObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return resourceGreenhouseDepartmentRead(ctx, d, meta)
}

func resourceGreenhouseDepartmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for departments."}}
}

func resourceGreenhouseDepartmentImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseDepartmentRead)
}

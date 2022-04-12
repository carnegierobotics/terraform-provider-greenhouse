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

func resourceGreenhouseOffice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseOfficeCreate,
		ReadContext:   resourceGreenhouseOfficeRead,
		UpdateContext: resourceGreenhouseOfficeUpdate,
		DeleteContext: resourceGreenhouseOfficeDelete,
		Exists:        resourceGreenhouseOfficeExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseOfficeImport,
		},
		Schema: schemaGreenhouseOffice(),
	}
}

func resourceGreenhouseOfficeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/offices/%d", id))
}

func resourceGreenhouseOfficeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.OfficeCreateInfo{
		Name:                 StringPtr(d.Get("name").(string)),
		Location:             StringPtr(d.Get("location.name").(string)),
		PrimaryContactUserId: IntPtr(d.Get("primary_contact_user_id").(int)),
		ParentId:             IntPtr(d.Get("parent_id").(int)),
	}
	id, err := greenhouse.CreateOffice(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseOfficeUpdate(ctx, d, meta)
}

func resourceGreenhouseOfficeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetOffice(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenOffice(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseOfficeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	updateObject := greenhouse.OfficeUpdateInfo{
		Name:     StringPtr(d.Get("name").(string)),
		Location: StringPtr(d.Get("location.name").(string)),
	}
	err = greenhouse.UpdateOffice(meta.(*greenhouse.Client), ctx, id, &updateObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return resourceGreenhouseOfficeRead(ctx, d, meta)
}

func resourceGreenhouseOfficeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for offices."}}
}

func resourceGreenhouseOfficeImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseOfficeRead)
}

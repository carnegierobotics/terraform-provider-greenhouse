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

func resourceGreenhouseUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseUserCreate,
		ReadContext:   resourceGreenhouseUserRead,
		UpdateContext: resourceGreenhouseUserUpdate,
		DeleteContext: resourceGreenhouseUserDelete,
		Exists:        resourceGreenhouseUserExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseUserImport,
		},
		Schema: schemaGreenhouseUser(),
	}
}

func resourceGreenhouseUserExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/users/%d", id))
}

func resourceGreenhouseUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var createObject greenhouse.UserCreateInfo
	if v, ok := d.Get("first_name").(string); ok && len(v) > 0 {
		createObject.FirstName = &v
	}
	if v, ok := d.Get("last_name").(string); ok && len(v) > 0 {
		createObject.LastName = &v
	}
	if v, ok := d.Get("primary_email_address").(string); ok && len(v) > 0 {
		createObject.Email = &v
	}
	if v, ok := d.Get("send_email").(bool); ok {
		createObject.SendEmail = &v
	}
	id, err := greenhouse.CreateUser(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseUserUpdate(ctx, d, meta)
}

func resourceGreenhouseUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetUser(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenUser(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if d.HasChange("disable_user") {
		if d.Get("disable_user").(bool) {
			err = greenhouse.DisableUser(meta.(*greenhouse.Client), ctx, id)
		} else {
			err = greenhouse.EnableUser(meta.(*greenhouse.Client), ctx, id)
		}
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	} else {
		updateObject := greenhouse.UserUpdateInfo{
			FirstName: StringPtr(d.Get("first_name").(string)),
			LastName:  StringPtr(d.Get("last_name").(string)),
		}
		err = greenhouse.UpdateUser(meta.(*greenhouse.Client), ctx, id, &updateObject)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return resourceGreenhouseUserRead(ctx, d, meta)
}

func resourceGreenhouseUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for users."}}
}

func resourceGreenhouseUserImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseUserRead)
}

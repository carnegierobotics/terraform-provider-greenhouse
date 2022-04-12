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
	"strconv"
)

func resourceGreenhouseCandidateTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseCandidateTagCreate,
		ReadContext:   resourceGreenhouseCandidateTagRead,
		UpdateContext: resourceGreenhouseCandidateTagUpdate,
		DeleteContext: resourceGreenhouseCandidateTagDelete,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseCandidateTagImport,
		},
		Schema: schemaGreenhouseTypeIdName(),
	}
}

func resourceGreenhouseCandidateTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.CandidateTag{
		Name: StringPtr(d.Get("name").(string)),
	}
	id, err := greenhouse.CreateCandidateTag(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseCandidateTagRead(ctx, d, meta)
}

func resourceGreenhouseCandidateTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetCandidateTag(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if obj == nil {
		tflog.Warn(ctx, fmt.Sprintf("Could not find tag with id %d", id))
		return nil
	}
	for k, v := range flattenCandidateTag(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseCandidateTagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for candidate_tag."}}
}

func resourceGreenhouseCandidateTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteCandidateTag(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return nil
}

func resourceGreenhouseCandidateTagImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseCandidateTagRead)
}

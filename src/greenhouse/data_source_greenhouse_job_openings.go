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

func dataSourceGreenhouseJobOpenings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseJobOpeningsRead,
		Schema: map[string]*schema.Schema{
			"job_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"openings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseJobOpening(),
				},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGreenhouseJobOpeningsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	jobId := d.Get("job_id").(int)
	status := d.Get("status").(string)
	list, err := greenhouse.GetAllJobOpenings(meta.(*greenhouse.Client), ctx, jobId, status)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	id := strconv.Itoa(jobId)
	if status != "" {
		id = fmt.Sprintf("%s-%s", id, status)
	}
	d.SetId(id)
	d.Set("openings", flattenJobOpenings(ctx, list))
	return nil
}

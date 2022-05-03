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
)

func dataSourceGreenhouseCandidates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCandidatesRead,
		Schema: map[string]*schema.Schema{
			"candidate_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"candidates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseCandidate(),
				},
			},
			"created_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"job_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGreenhouseCandidatesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllCandidates(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	nameList := make([]string, 0, len(*list))
	for _, item := range *list {
		var name string
		if v := item.FirstName; v != nil {
			name += *v
		}
		if v := item.LastName; v != nil {
			name += *v
		}
		if name != "" {
			nameList = append(nameList, name)
		}
	}
	d.SetId("all")
	d.Set("candidates", flattenCandidates(ctx, list))
	d.Set("names", nameList)
	return nil
}

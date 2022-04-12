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

func dataSourceGreenhouseCandidateTags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCandidateTagsRead,
		Schema:      schemaGreenhouseCandidateTags(),
	}
}

func dataSourceGreenhouseCandidateTagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	candidateId, ok := d.GetOk("candidate_id")
	var id string
	var tags *[]greenhouse.CandidateTag
	var err error
	if ok {
		id = strconv.Itoa(candidateId.(int))
		tags, err = greenhouse.GetTagsForCandidate(meta.(*greenhouse.Client), ctx, candidateId.(int))
	} else {
		id = "all"
		tags, err = greenhouse.GetAllCandidateTags(meta.(*greenhouse.Client), ctx)
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(id)
	d.Set("tags", flattenCandidateTags(ctx, tags))
	return nil
}

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

func dataSourceGreenhouseActivityFeed() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseActivityFeedRead,
		Schema:      schemaGreenhouseActivityFeed(),
	}
}

func dataSourceGreenhouseActivityFeedRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	candidateId := d.Get("candidate_id").(int)
	obj, err := greenhouse.GetActivityFeed(meta.(*greenhouse.Client), ctx, candidateId)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(candidateId))
	d.Set("activities", flattenActivities(ctx, &obj.Activities))
	d.Set("emails", flattenEmails(ctx, &obj.Emails))
	d.Set("notes", flattenNotes(ctx, &obj.Notes))
	return nil
}

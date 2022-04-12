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

func dataSourceGreenhouseTrackingLink() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseTrackingLinkRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseTrackingLinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	token := d.Get("token").(string)
	link, err := greenhouse.GetTrackingLinkData(meta.(*greenhouse.Client), ctx, token)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(*link.Id))
	d.Set("created_at", link.CreatedAt)
	d.Set("credited_to", flattenUser(ctx, link.CreditedTo))
	d.Set("job_board", flattenJobBoard(ctx, link.JobBoard))
	d.Set("job_id", link.JobId)
	d.Set("job_post_id", link.JobPostId)
	d.Set("related_post_id", link.RelatedPostId)
	d.Set("related_post_type", link.RelatedPostType)
	d.Set("source", flattenSource(ctx, link.Source))
	d.Set("updated_at", link.UpdatedAt)
	return nil
}

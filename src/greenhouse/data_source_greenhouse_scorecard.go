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

func dataSourceGreenhouseScorecard() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseScorecardRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseScorecardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("id").(int)
	card, err := greenhouse.GetScorecard(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(*card.Id))
	d.Set("application_id", card.ApplicationId)
	d.Set("attributes", flattenScorecardAttributes(ctx, &card.Attributes))
	d.Set("candidate_id", card.CandidateId)
	d.Set("created_at", card.CreatedAt)
	d.Set("interview", card.Interview)
	convertedStep := greenhouse.TypeIdName(*card.InterviewStep)
	d.Set("interview_step", flattenTypeIdName(ctx, &convertedStep))
	d.Set("interviewer", flattenUser(ctx, card.Interviewer))
	d.Set("overall_recommendation", card.OverallRecommendation)
	d.Set("questions", flattenScorecardQuestions(ctx, &card.Questions))
	d.Set("ratings", card.Ratings)
	d.Set("submitted_at", card.SubmittedAt)
	d.Set("submitted_by", flattenUser(ctx, card.SubmittedBy))
	d.Set("updated_at", card.UpdatedAt)
	return nil
}

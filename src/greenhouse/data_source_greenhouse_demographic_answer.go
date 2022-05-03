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

func dataSourceGreenhouseDemographicAnswer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDemographicAnswerRead,
		Schema:      schemaGreenhouseDemographicAnswer(),
	}
}

func dataSourceGreenhouseDemographicAnswerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("id").(int)
	answer, err := greenhouse.GetDemographicAnswer(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(*answer.Id))
	d.Set("application_id", answer.ApplicationId)
	d.Set("created_at", answer.CreatedAt)
	d.Set("demographic_answer_option_id", answer.DemographicAnswerOptionId)
	d.Set("demographic_question_id", answer.DemographicQuestionId)
	d.Set("free_form_text", answer.FreeFormText)
	d.Set("updated_at", answer.UpdatedAt)
	return nil
}

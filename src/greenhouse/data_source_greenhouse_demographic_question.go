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

func dataSourceGreenhouseDemographicQuestion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDemographicQuestionRead,
		Schema:      schemaGreenhouseDemographicQuestion(),
	}
}

func dataSourceGreenhouseDemographicQuestionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	list, err := greenhouse.GetAllDemographicQuestions(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for _, question := range *list {
		if *question.Name == name {
			d.SetId(strconv.Itoa(*question.Id))
			d.Set("active", question.Active)
			d.Set("answer_type", question.AnswerType)
			d.Set("demographic_question_set_id", question.DemographicQuestionSetId)
			d.Set("required", question.Required)
			d.Set("translations", flattenTranslations(ctx, &question.Translations))
			return nil
		}
	}
	return nil
}

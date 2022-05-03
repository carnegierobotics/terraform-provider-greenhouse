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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDemographicQuestion() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"answer_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"demographic_question_set_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"required": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"translations": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTranslation(),
			},
		},
	}
}

func flattenDemographicQuestion(ctx context.Context, item *greenhouse.DemographicQuestion) map[string]interface{} {
	question := make(map[string]interface{})
	if v := item.Active; v != nil {
		question["active"] = *v
	}
	if v := item.AnswerType; v != nil {
		question["answer_type"] = *v
	}
	if v := item.DemographicQuestionSetId; v != nil {
		question["demographic_question_set_id"] = *v
	}
	if v := item.Name; v != nil {
		question["name"] = *v
	}
	if v := item.Required; v != nil {
		question["required"] = *v
	}
	if v := item.Translations; len(v) > 0 {
		question["translations"] = flattenTranslations(ctx, &v)
	}
	return question
}

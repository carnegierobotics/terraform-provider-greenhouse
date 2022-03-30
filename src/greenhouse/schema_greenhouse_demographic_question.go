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
	question["active"] = item.Active
	question["answer_type"] = item.AnswerType
	question["demographic_question_set_id"] = item.DemographicQuestionSetId
	question["name"] = item.Name
	question["required"] = item.Required
	question["translations"] = flattenTranslations(ctx, &item.Translations)
	return question
}

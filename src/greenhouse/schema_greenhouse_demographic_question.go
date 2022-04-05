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

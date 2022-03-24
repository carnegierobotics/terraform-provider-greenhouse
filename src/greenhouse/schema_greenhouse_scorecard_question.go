package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScorecardQuestion() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"answer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"question": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func flattenScorecardQuestions(ctx context.Context, list *[]greenhouse.ScorecardQuestion) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenScorecardQuestion(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenScorecardQuestion(ctx context.Context, item *greenhouse.ScorecardQuestion) map[string]interface{} {
	question := make(map[string]interface{})
	question["answer"] = item.Answer
	question["question"] = item.Question
	return question
}

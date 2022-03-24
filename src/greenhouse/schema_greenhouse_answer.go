package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseAnswer() map[string]*schema.Schema {
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

func flattenAnswers(ctx context.Context, list *[]greenhouse.Answer) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenAnswer(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenAnswer(ctx context.Context, item *greenhouse.Answer) map[string]interface{} {
	answer := make(map[string]interface{})
	answer["question"] = item.Question
	answer["answer"] = item.Answer
	return answer
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterviewQuestion() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"question": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func flattenInterviewQuestions(ctx context.Context, list *[]greenhouse.InterviewQuestion) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenInterviewQuestion(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenInterviewQuestion(ctx context.Context, item *greenhouse.InterviewQuestion) map[string]interface{} {
	question := make(map[string]interface{})
	question["question"] = item.Question
	return question
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateInterviewQuestions(ctx context.Context, source *[]interface{}) (*[]greenhouse.InterviewQuestion, diag.Diagnostics) {
	list := make([]greenhouse.InterviewQuestion, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateInterviewQuestion(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateInterviewQuestion(ctx context.Context, source *map[string]interface{}) (*greenhouse.InterviewQuestion, diag.Diagnostics) {
	var obj greenhouse.InterviewQuestion
	if v, ok := (*source)["question"].(string); ok && len(v) > 0 {
		obj.Question = &v
	}
	return &obj, nil
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
	if v := item.Question; v != nil {
		question["question"] = *v
	}
	return question
}

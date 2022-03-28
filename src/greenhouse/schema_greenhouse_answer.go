package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func inflateAnswers(ctx context.Context, source interface{}) *[]greenhouse.Answer {
	var list []greenhouse.Answer
	convertType(ctx, source, list)
	return &list
}

func inflateAnswer(ctx context.Context, source map[string]interface{}) *greenhouse.Answer {
	var item greenhouse.Answer
	convertType(ctx, source, item)
	return &item
}

func flattenAnswers(ctx context.Context, list *[]greenhouse.Answer) []interface{} {
	if list != nil {
		tflog.Debug(ctx, "Flattening answers.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenAnswer(ctx, &item)
		}
		tflog.Debug(ctx, "Finished flattening answers.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenAnswer(ctx context.Context, item *greenhouse.Answer) map[string]interface{} {
	tflog.Debug(ctx, "Flattening one answer.")
	answer := make(map[string]interface{})
	answer["question"] = item.Question
	answer["answer"] = item.Answer
	tflog.Debug(ctx, "Finished flattening answer.")
	return answer
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateAnswers(ctx context.Context, source *[]interface{}) (*[]greenhouse.Answer, diag.Diagnostics) {
	list := make([]greenhouse.Answer, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateAnswer(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateAnswer(ctx context.Context, source *map[string]interface{}) (*greenhouse.Answer, diag.Diagnostics) {
	var obj greenhouse.Answer
	if v, ok := (*source)["answer"].(string); ok && len(v) > 0 {
		obj.Answer = &v
	}
	if v, ok := (*source)["question"].(string); ok && len(v) > 0 {
		obj.Question = &v
	}
	return &obj, nil
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
	if v := item.Question; v != nil {
		answer["question"] = *v
	}
	if v := item.Answer; v != nil {
		answer["answer"] = *v
	}
	tflog.Debug(ctx, "Finished flattening answer.")
	return answer
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterviewKit() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Type:     schema.TypeString,
			Required: true,
		},
		"questions": {
			Type:     schema.TypeSet,
			Required: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseInterviewQuestion(),
			},
		},
	}
}

func inflateInterviewKits(ctx context.Context, source *[]interface{}) (*[]greenhouse.InterviewKit, diag.Diagnostics) {
	list := make([]greenhouse.InterviewKit, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateInterviewKit(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateInterviewKit(ctx context.Context, source *map[string]interface{}) (*greenhouse.InterviewKit, diag.Diagnostics) {
	var obj greenhouse.InterviewKit
	if v, ok := (*source)["content"].(string); ok && len(v) > 0 {
		obj.Content = v
	}
	if v, ok := (*source)["questions"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateInterviewQuestions(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Questions = *list
	}
	return &obj, nil
}

func flattenInterviewKit(ctx context.Context, item *greenhouse.InterviewKit) []interface{} {
	kit := make([]interface{}, 1, 1)
	oneKit := make(map[string]interface{})
	oneKit["content"] = item.Content
	oneKit["questions"] = flattenInterviewQuestions(ctx, &item.Questions)
	kit[0] = oneKit
	return kit
}

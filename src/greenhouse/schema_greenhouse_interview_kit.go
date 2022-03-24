package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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

func flattenInterviewKit(ctx context.Context, item *greenhouse.InterviewKit) []interface{} {
	kit := make([]interface{}, 1, 1)
	oneKit := make(map[string]interface{})
	oneKit["content"] = item.Content
	oneKit["questions"] = flattenInterviewQuestions(ctx, &item.Questions)
	kit[0] = oneKit
	return kit
}

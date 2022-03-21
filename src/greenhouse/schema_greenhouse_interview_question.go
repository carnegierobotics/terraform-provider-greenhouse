package greenhouse

import (
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

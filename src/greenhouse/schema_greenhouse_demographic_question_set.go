package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDemographicQuestionSet() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"title": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

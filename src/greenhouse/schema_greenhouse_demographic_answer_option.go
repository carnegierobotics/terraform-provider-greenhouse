package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDemographicAnswerOption() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"demographic_question_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"free_form": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"translations": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTranslation(),
			},
		},
	}
}

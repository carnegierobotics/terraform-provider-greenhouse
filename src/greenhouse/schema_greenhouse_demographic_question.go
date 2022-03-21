package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseDemographicQuestion() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"answer_type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"demographic_question_set_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"required": {
			Type:     schema.TypeBool,
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

package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseCustomFieldOption() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Description: "The name of the new custom field option.",
			Required:    true,
		},
		"priority": {
			Type:        schema.TypeInt,
			Description: "Numeric value for ordering the custom field options.",
			Required:    true,
		},
		"external_id": {
			Type:        schema.TypeString,
			Description: "The external_id for the custom field.",
			Optional:    true,
		},
	}
}

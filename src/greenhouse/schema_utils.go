package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseValueType() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"value": {
			Type:     schema.TypeString,
			Optional: false,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: false,
		},
	}
}

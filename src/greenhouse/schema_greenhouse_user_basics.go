package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUserBasics() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
	}
}

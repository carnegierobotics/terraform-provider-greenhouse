package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseRecruiter() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"email": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

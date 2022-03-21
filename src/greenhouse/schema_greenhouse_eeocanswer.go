package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEEOCAnswer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"message": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

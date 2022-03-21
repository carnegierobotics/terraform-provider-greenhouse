package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJobBoard() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"company_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"url_token": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

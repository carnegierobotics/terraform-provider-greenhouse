package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func schemaGreenhouseHiringTeam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

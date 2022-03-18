package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScorecardQuestion() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"answer": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"question": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

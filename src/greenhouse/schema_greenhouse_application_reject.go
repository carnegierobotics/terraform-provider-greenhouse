package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApplicationReject() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"notes": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"rejection_email": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRejectionEmail(),
			},
		},
		"rejection_reason": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseRejectionEmail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_template_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"send_email_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

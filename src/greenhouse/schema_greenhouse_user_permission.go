package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUserPermission() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"user_role_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

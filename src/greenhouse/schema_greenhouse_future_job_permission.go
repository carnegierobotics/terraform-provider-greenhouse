package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseFutureJobPermission() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"department_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"external_department_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"external_office_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"office_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
    "user_role_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
	}
}

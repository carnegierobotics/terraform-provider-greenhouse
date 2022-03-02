package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func schemaGreenhouseOffice() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"location": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"primary_contact_user_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"child_ids": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	}
}

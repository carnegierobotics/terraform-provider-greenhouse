package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUser() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"primary_email_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"send_email": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"disable_user": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"disabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"site_admin": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"emails": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"linked_candidate_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	}
}

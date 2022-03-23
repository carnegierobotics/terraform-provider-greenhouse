package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmailTemplate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"body": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"cc": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"default": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"from": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"html_body": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"user": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUser(),
			},
		},
	}
}

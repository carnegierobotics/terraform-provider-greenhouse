package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
    "subject": {
      Type: schema.TypeString,
      Computed: true,
    },
    "body": {
      Type:     schema.TypeString,
      Computed: true,
    },
    "to": {
      Type:   schema.TypeString,
      Computed: true,
    },
    "from": {
      Type: schema.TypeString,
      Computed: true,
    },
    "cc": {
      Type: schema.TypeString,
      Computed: true,
    },
    "user": {
      Type:     schema.TypeSet,
      MaxItems: 1,
      Computed: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseUserBasics(),
      },
    },
	}
}

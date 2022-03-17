package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseNote() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
    "body": {
      Type:     schema.TypeString,
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
    "private": {
      Type: schema.TypeBool,
      Computed: true,
    },
    "visiblity": {
      Type: schema.TypeString,
      Computed: true,
    },
    "visibility": {
      Type: schema.TypeString,
      Computed: true,
    },
	}
}

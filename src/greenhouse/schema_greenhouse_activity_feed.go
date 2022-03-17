package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseActivityFeed() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"notes": {
			Type:     schema.TypeList,
			Computed: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseNote(),
      },
		},
		"emails": {
			Type:     schema.TypeList,
			Computed: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseEmail(),
      },
		},
		"activities": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseActivity(),
			},
		},
	}
}

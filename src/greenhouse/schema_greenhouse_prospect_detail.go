package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseProspectDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "prospect_owner": {
      Type: schema.TypeString,
      Optional: true,
    },
    "prospect_pool": {
      Type: schema.TypeString,
      Optional: true,
    },
    "prospect_stage": {
      Type: schema.TypeString,
      Optional: true,
    },
	}
}

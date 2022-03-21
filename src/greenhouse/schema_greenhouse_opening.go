package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOpening() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "custom_fields": {
      Type: schema.TypeSet,
      Optional: true,
      Elem: schema.Schema{
        Type: schema.TypeMap,
        Elem: &schema.Schema{
          Type: schema.TypeString,
        },
      }
    },
	}
}

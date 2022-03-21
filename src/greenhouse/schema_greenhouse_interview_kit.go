package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterviewKit() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "content": {
      Type: schema.TypeString,
      Required: true,
    },
    "questions": {
      Type: schema.TypeSet,
      Required: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseInterviewQuestion(),
      },
    },
	}
}

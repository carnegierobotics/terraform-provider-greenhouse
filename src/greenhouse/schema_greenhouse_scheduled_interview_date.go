package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseScheduledInterviewDate() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "date": {
      Type: schema.TypeString,
      Optional: true,
    },
    "date_time": {
      Type: schema.TypeString,
      Optional: true,
    },
	}
}

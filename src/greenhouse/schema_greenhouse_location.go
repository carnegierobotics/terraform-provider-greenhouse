package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseLocation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func flattenLocation(ctx context.Context, item *greenhouse.Location) []interface{} {
	location := make([]interface{}, 1, 1)
	oneLocation := make(map[string]interface{})
	if item.Name != "" {
		oneLocation["name"] = item.Name
	}
	location[0] = oneLocation
	return location
}

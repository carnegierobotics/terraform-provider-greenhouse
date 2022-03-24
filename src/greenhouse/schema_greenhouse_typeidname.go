package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseTypeIdName() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func flattenTypeIdName(ctx context.Context, item *greenhouse.TypeIdName) map[string]interface{} {
	newItem := make(map[string]interface{})
	newItem["name"] = item.Name
	return newItem
}

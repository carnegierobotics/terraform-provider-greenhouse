package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func schemaGreenhouseTypeIdName() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateTypeIdName(ctx context.Context, source interface{}) *greenhouse.TypeIdName {
	var item greenhouse.TypeIdName
	convertType(ctx, source, item)
	return &item
}

func flattenTypeIdName(ctx context.Context, item *greenhouse.TypeIdName) map[string]interface{} {
	newItem := make(map[string]interface{})
	newItem["id"] = strconv.Itoa(item.Id)
	newItem["name"] = item.Name
	return newItem
}

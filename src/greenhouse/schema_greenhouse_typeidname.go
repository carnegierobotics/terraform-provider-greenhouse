package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func schemaGreenhouseTypeIdName() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateTypesIdName(ctx context.Context, source *[]interface{}) (*[]greenhouse.TypeIdName, diag.Diagnostics) {
	list := make([]greenhouse.TypeIdName, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateTypeIdName(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateTypeIdName(ctx context.Context, source *map[string]interface{}) (*greenhouse.TypeIdName, diag.Diagnostics) {
	var item greenhouse.TypeIdName
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		item.Name = &v
	}
	return &item, nil
}

func flattenTypeIdName(ctx context.Context, item *greenhouse.TypeIdName) map[string]interface{} {
	newItem := make(map[string]interface{})
	if v := item.Id; v != nil {
		newItem["id"] = strconv.Itoa(*v)
	}
	if v := item.Name; v != nil {
		newItem["name"] = *v
	}
	return newItem
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseSource() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func inflateSource(ctx context.Context, source interface{}) *greenhouse.Source {
	var item greenhouse.Source
	convertType(ctx, source, item)
	return &item
}

func flattenSource(ctx context.Context, item *greenhouse.Source) map[string]interface{} {
	source := make(map[string]interface{})
	if v := item.PublicName; v != nil {
		source["public_name"] = v
	}
	if v := item.Type; v != nil {
		source["type"] = flattenTypeIdName(ctx, v)
	}
	return source
}

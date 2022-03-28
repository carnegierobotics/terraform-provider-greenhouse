package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseRejectionReason() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"include_defaults": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"per_page": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"type": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func inflateRejectionReason(ctx context.Context, source interface{}) *greenhouse.RejectionReason {
	var item greenhouse.RejectionReason
	convertType(ctx, source, item)
	return &item
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseProspectDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prospect_owner": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"prospect_pool": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
		"prospect_stage": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseTypeIdName(),
			},
		},
	}
}

func flattenProspectDetail(ctx context.Context, item *greenhouse.ProspectDetail) map[string]interface{} {
	detail := make(map[string]interface{})
	detail["prospect_owner"] = flattenTypeIdName(ctx, &item.ProspectOwner)
	detail["prospect_pool"] = flattenTypeIdName(ctx, &item.ProspectPool)
	detail["prospect_stage"] = flattenTypeIdName(ctx, &item.ProspectStage)
	return detail
}

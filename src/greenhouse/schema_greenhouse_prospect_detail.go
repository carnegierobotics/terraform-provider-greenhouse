package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseProspectDetail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"prospect_owner": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"prospect_pool": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"prospect_stage": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func flattenProspectDetail(ctx context.Context, item *greenhouse.ProspectDetail) map[string]interface{} {
	detail := make(map[string]interface{})
	detail["prospect_owner"] = item.ProspectOwner
	detail["prospect_pool"] = item.ProspectPool
	detail["prospect_stage"] = item.ProspectStage
	return detail
}

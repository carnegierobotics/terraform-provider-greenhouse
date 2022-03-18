package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApproverGroup() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"approvals_required": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"approvers": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApprover(),
			},
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"offer_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"resolved_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

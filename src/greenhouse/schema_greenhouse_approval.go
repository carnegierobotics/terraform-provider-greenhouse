package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApproval() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"approval_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"approval_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"approver_groups": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApproverGroup(),
			},
		},
		"job_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"offer_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"requested_by_user_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"sequential": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

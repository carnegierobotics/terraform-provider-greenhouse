package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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

func flattenApproverGroups(ctx context.Context, list *[]greenhouse.ApproverGroup) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApproverGroup(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApproverGroup(ctx context.Context, item *greenhouse.ApproverGroup) map[string]interface{} {
	approver := make(map[string]interface{})
	approver["approvals_required"] = item.ApprovalsRequired
	approver["approvers"] = flattenApprovers(ctx, &item.Approvers)
	approver["created_at"] = item.CreatedAt
	approver["job_id"] = item.JobId
	approver["offer_id"] = item.OfferId
	approver["priority"] = item.Priority
	approver["resolved_at"] = item.ResolvedAt
	return approver
}

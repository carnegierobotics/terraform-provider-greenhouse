package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

func flattenApprovals(ctx context.Context, list *[]greenhouse.Approval) []interface{} {
	if list != nil {
		tflog.Debug(ctx, "Flattening approvals.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApproval(ctx, &item)
		}
		tflog.Debug(ctx, "Finished flattening approvals.")
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApproval(ctx context.Context, item *greenhouse.Approval) map[string]interface{} {
	approval := make(map[string]interface{})
	approval["approval_status"] = item.ApprovalStatus
	approval["approval_type"] = item.ApprovalType
	approval["approver_groups"] = flattenApproverGroups(ctx, &item.ApproverGroups)
	approval["job_id"] = item.JobId
	approval["offer_id"] = item.OfferId
	approval["requested_by_user_id"] = item.RequestedByUserId
	approval["sequential"] = item.Sequential
	approval["version"] = item.Version
	return approval
}

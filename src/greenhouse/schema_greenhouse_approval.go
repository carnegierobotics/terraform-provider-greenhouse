package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			Required: true,
		},
		"approver_groups": {
			Type:     schema.TypeList,
      Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseApproverGroup(),
			},
		},
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"offer_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
    "request_approval": {
      Type: schema.TypeBool,
      Optional: true,
    },
		"requested_by_user_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"sequential": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func inflateApprovals(ctx context.Context, source *[]interface{}) (*[]greenhouse.Approval, diag.Diagnostics) {
	list := make([]greenhouse.Approval, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApproval(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApproval(ctx context.Context, source *map[string]interface{}) (*greenhouse.Approval, diag.Diagnostics) {
	var obj greenhouse.Approval
	if v, ok := (*source)["approval_status"].(string); ok && len(v) > 0 {
		obj.ApprovalStatus = v
	}
	if v, ok := (*source)["approval_type"].(string); ok && len(v) > 0 {
		obj.ApprovalType = v
	}
	if v, ok := (*source)["approver_groups"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateApproverGroups(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.ApproverGroups = *list
	}
	if v, ok := (*source)["job_id"].(int); ok {
		obj.JobId = v
	}
	if v, ok := (*source)["offer_id"].(int); ok {
		obj.OfferId = v
	}
	if v, ok := (*source)["requested_by_user_id"].(int); ok {
		obj.RequestedByUserId = v
	}
	if v, ok := (*source)["sequential"].(bool); ok {
		obj.Sequential = v
	}
	if v, ok := (*source)["version"].(int); ok {
		obj.Version = v
	}
	return &obj, nil
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

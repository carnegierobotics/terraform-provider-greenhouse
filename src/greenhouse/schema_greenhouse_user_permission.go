package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUserPermission() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"user_role_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func flattenUserPermissions(ctx context.Context, list *[]greenhouse.UserPermission) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenUserPermission(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenUserPermission(ctx context.Context, item *greenhouse.UserPermission) map[string]interface{} {
	permission := make(map[string]interface{})
	permission["job_id"] = item.JobId
	permission["user_role_id"] = item.UserRoleId
	return permission
}

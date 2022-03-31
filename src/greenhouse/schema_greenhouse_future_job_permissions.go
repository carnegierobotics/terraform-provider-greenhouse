package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseFutureJobPermissionList() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"permissions": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseFutureJobPermission(),
			},
		},
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func flattenFutureJobPermissionsList(ctx context.Context, list *[]greenhouse.FutureJobPermission) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenFutureJobPermission(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

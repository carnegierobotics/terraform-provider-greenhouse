package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUserBasics() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func flattenUsersBasics(ctx context.Context, list *[]greenhouse.User) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenUserBasics(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenUserBasics(ctx context.Context, item *greenhouse.User) map[string]interface{} {
	tflog.Debug(ctx, "Flattening one user basics.")
	user := make(map[string]interface{})
	if v := item.EmployeeId; v != nil {
		user["employee_id"] = *v
	}
	if v := item.FirstName; v != nil {
		user["first_name"] = *v
	}
	if v := item.LastName; v != nil {
		user["last_name"] = *v
	}
	if v := item.Name; v != nil {
		user["name"] = *v
	}
	tflog.Debug(ctx, "Finished flattening user basics.")
	return user
}

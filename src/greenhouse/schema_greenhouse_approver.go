package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApprover() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_addresses": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"employee_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func flattenApprovers(ctx context.Context, list *[]greenhouse.Approver) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenApprover(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenApprover(ctx context.Context, item *greenhouse.Approver) map[string]interface{} {
	approver := make(map[string]interface{})
	approver["email_addresses"] = item.EmailAddresses
	approver["employee_id"] = item.EmployeeId
	approver["name"] = item.Name
	return approver
}

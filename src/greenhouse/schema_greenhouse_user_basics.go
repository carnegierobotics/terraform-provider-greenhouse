package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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
			Default:  "",
		},
	}
}

func flattenUserBasics(ctx context.Context, item *greenhouse.User) map[string]interface{} {
	user := make(map[string]interface{})
	user["employee_id"] = item.EmployeeId
	user["first_name"] = item.FirstName
	user["last_name"] = item.LastName
	user["name"] = item.Name
	return user
}

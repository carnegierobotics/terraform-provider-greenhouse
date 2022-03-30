package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseHiringMember() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
    /*
		"active": {
			Type:     schema.TypeBool,
			Computed: true,
		},
    */
		"responsible": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"responsible_for_future_candidates": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"responsible_for_active_candidates": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"responsible_for_inactive_candidates": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"employee_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func flattenHiringTeamMember(ctx context.Context, item greenhouse.HiringMember) (map[string]interface{}, error) {
	tflog.Debug(ctx, "User data", "user", fmt.Sprintf("%+v", item))
	member := make(map[string]interface{})
	member["user_id"] = item.Id
	member["active"] = item.Active
	member["first_name"] = item.FirstName
	member["last_name"] = item.LastName
	member["name"] = item.Name
	member["responsible"] = item.Responsible
	member["employee_id"] = item.EmployeeId
	return member, nil
}

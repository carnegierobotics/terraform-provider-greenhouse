package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseHiringMember() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"active": {
			Type:     schema.TypeBool,
			Computed: true,
		},
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

// Hiring team is map[string][]HiringMember
func flattenHiringTeam(list *[]greenhouse.HiringMember) map[string]interface{} {
  if list != nil {
    flatMap := make(map[string]interface{})
    for k, team := range *list {
      flatMap[k] := make([]interface{}, len(team), len(team))
      for i, item := range *list {
        member := make(map[string]interface{})
        member["user_id"] = item.UserId
        member["active"] = item.Active
        member["first_name"] = item.FirstName
        member["last_name"] = item.LastName
        member["name"] = item.Name
        member["responsible"] = item.Responsible
        member["employee_id"] = item.EmployeeId
        flatMap[k][i] = member
      }
    }
    return flatMap
  }
  return make([]interface{}, 0)
}

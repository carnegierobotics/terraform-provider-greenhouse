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
			Optional: true,
		},
		"active": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"responsible": {
			Type:     schema.TypeBool,
      Optional: true,
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
      Optional: true,
			Computed: true,
		},
	}
}

/*
Hiring team is map[string][]HiringMember
{
  "team1": [
    {
      "name": "member 1",
      "first_name": "member"
    },
    {
      "name": "member 2",
      "first_name": "member"
    }
  ],
  "team2": [
    {
      "name": "member 3",
      "first_name": "member"
    },
  ]
}
*/
func flattenHiringTeam(ctx context.Context, list *map[string][]greenhouse.HiringMember) []interface{} {
	if list != nil {
		flatMap := make([]interface{}, 1)
    flatMap[0] = make(map[string]interface{})
		for k, team := range *list {
			flatMap[0].(map[string]interface{})[k] = make([]interface{}, len(team), len(team))
      if team != nil {
			  for i, item := range team {
				  member := make(map[string]interface{})
				  member["user_id"] = item.UserId
				  member["active"] = item.Active
				  member["first_name"] = item.FirstName
				  member["last_name"] = item.LastName
				  member["name"] = item.Name
				  member["responsible"] = item.Responsible
				  member["employee_id"] = item.EmployeeId
				  flatMap[0].(map[string]interface{})[k].([]interface{})[i] = member
        }
			}
		}
    tflog.Debug(ctx, "Flattened team", "team", fmt.Sprintf("%+v", flatMap))
		return flatMap
	}
	return make([]interface{}, 0)
}

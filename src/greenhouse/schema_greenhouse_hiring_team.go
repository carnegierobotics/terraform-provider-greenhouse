package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseHiringTeam() map[string]*schema.Schema {
  return map[string]*schema.Schema{
    "name": {
      Type:     schema.TypeString,
      Required: true,
    },
    "members": {
      Type:     schema.TypeList,
      Required: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseHiringMember(),
      },
    },
  }
}

func schemaGreenhouseHiringMember() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"user_id": {
			Type:     schema.TypeInt,
			Required: true,
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
			Optional: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Optional: true,
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
    flatMap := make([]interface{}, len(*list), len(*list))
    teamCount := 0
    for k, team := range *list {
      flatTeam := make(map[string]interface{})
      flatTeam["name"] = k
      flatTeam["members"] = flattenOneTeam(ctx, team)
      flatMap[teamCount] = flatTeam
      teamCount++
    }
    tflog.Debug(ctx, "Flattened hiring team", fmt.Sprintf("%+v", flatMap))
    return flatMap
  }
  return make([]interface{}, 0)
}

func flattenOneTeam(ctx context.Context, team []greenhouse.HiringMember) []interface{} {
	if team != nil {
    flatMap := make([]interface{}, len(team), len(team))
	  for i, member := range team {
      member, _ := flattenHiringTeamMember(ctx, member)
  	  flatMap[i] = member
		}
	  tflog.Debug(ctx, "Flattened team", "team", fmt.Sprintf("%+v", flatMap))
	  return flatMap
	}
	return make([]interface{}, 0)
}

func flattenHiringTeamMember(ctx context.Context, item greenhouse.HiringMember) (map[string]interface{}, error) {
	tflog.Debug(ctx, "User data", "user", fmt.Sprintf("%+v", item))
	member := make(map[string]interface{})
	member["user_id"] = item.UserId
	member["active"] = item.Active
	member["first_name"] = item.FirstName
	member["last_name"] = item.LastName
	member["name"] = item.Name
	member["responsible"] = item.Responsible
	member["employee_id"] = item.EmployeeId
  return member, nil
}

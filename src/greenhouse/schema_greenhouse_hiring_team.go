package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func inflateHiringTeams(ctx context.Context, source *map[string]interface{}) (*map[string][]greenhouse.HiringMember, diag.Diagnostics) {
	newMap := make(map[string][]greenhouse.HiringMember)
	for k, v := range *source {
		team := v.([]interface{})
		inflatedTeam, err := inflateHiringTeam(ctx, &team)
		if err != nil {
			return nil, err
		}
		newMap[k] = *inflatedTeam
	}
	return &newMap, nil
}

func inflateHiringTeam(ctx context.Context, source *[]interface{}) (*[]greenhouse.HiringMember, diag.Diagnostics) {
	list := make([]greenhouse.HiringMember, len(*source), len(*source))
	for i, item := range *source {
		n := item.(map[string]interface{})
		member, err := inflateHiringTeamMember(ctx, &n)
		if err != nil {
			return nil, err
		}
		list[i] = *member
	}
	return &list, nil
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

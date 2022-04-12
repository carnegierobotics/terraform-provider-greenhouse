/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of terraform-provider-greenhouse.

terraform-provider-greenhouse is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

terraform-provider-greenhouse is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with terraform-provider-greenhouse. If not, see <https://www.gnu.org/licenses/>.
*/
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
		"job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"teams": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseHiringSubTeam(),
			},
		},
	}
}

func schemaGreenhouseHiringSubTeam() map[string]*schema.Schema {
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

func inflateHiringSubteams(ctx context.Context, source *map[string]interface{}) (*map[string][]greenhouse.HiringMember, diag.Diagnostics) {
	newMap := make(map[string][]greenhouse.HiringMember)
	for k, v := range *source {
		team := v.([]interface{})
		inflatedTeam, err := inflateHiringSubteam(ctx, &team)
		if err != nil {
			return nil, err
		}
		newMap[k] = *inflatedTeam
	}
	return &newMap, nil
}

func inflateHiringSubteam(ctx context.Context, source *[]interface{}) (*[]greenhouse.HiringMember, diag.Diagnostics) {
	list := make([]greenhouse.HiringMember, len(*source), len(*source))
	for i, item := range *source {
		n := item.(map[string]interface{})
		member, err := inflateHiringSubteamMember(ctx, &n)
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
func flattenHiringSubteams(ctx context.Context, list *map[string][]greenhouse.HiringMember) []interface{} {
	if list != nil {
		flatMap := make([]interface{}, 0, 0)
		for k, team := range *list {
			flatTeam := make(map[string]interface{})
			if v := flattenOneTeam(ctx, team); len(v) > 0 {
				flatTeam["name"] = k
				flatTeam["members"] = flattenOneTeam(ctx, team)
				flatMap = append(flatMap, flatTeam)
			}
		}
		tflog.Trace(ctx, "Flattened hiring team", fmt.Sprintf("%+v", flatMap))
		return flatMap
	}
	return make([]interface{}, 0)
}

func flattenOneTeam(ctx context.Context, team []greenhouse.HiringMember) []interface{} {
	if team != nil {
		flatMap := make([]interface{}, len(team), len(team))
		for i, member := range team {
			member, _ := flattenHiringSubteamMember(ctx, member)
			flatMap[i] = member
		}
		tflog.Trace(ctx, "Flattened team", "team", fmt.Sprintf("%+v", flatMap))
		return flatMap
	}
	return make([]interface{}, 0)
}

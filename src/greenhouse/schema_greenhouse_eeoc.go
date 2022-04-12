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
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func schemaGreenhouseEEOC() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"candidate_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"disability_status": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
		"gender": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
		"race": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
		"submitted_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"veteran_status": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEEOCAnswer(),
			},
		},
	}
}

func flattenEEOCAnswer(ctx context.Context, item *greenhouse.EEOCAnswer) map[string]interface{} {
	answer := make(map[string]interface{})
	if v := item.Description; v != nil {
		answer["description"] = *v
	}
	if v := item.Id; v != nil {
		answer["id"] = strconv.Itoa(*v)
	}
	if v := item.Message; v != nil {
		answer["message"] = *v
	}
	return answer
}

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
)

func schemaGreenhouseOffer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"application_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"candidate_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"keyed_custom_fields": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseKeyedCustomField(),
				},
			},
		},
		"opening": {
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseJobOpening(),
			},
		},
		"resolved_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"sent_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"starts_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"version": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func flattenOffer(ctx context.Context, item *greenhouse.Offer) map[string]interface{} {
	offer := make(map[string]interface{})
	if v := item.ApplicationId; v != nil {
		offer["application_id"] = *v
	}
	if v := item.CandidateId; v != nil {
		offer["candidate_id"] = *v
	}
	if v := item.CreatedAt; v != nil {
		offer["created_at"] = *v
	}
	if v := item.CustomFields; v != nil {
		offer["custom_fields"] = v
	}
	if v := item.JobId; v != nil {
		offer["job_id"] = *v
	}
	if v := item.KeyedCustomFields; len(v) > 0 {
		offer["keyed_custom_fields"] = v
	}
	if v := item.Opening; v != nil {
		offer["opening"] = flattenJobOpening(ctx, v)
	}
	if v := item.ResolvedAt; v != nil {
		offer["resolved_at"] = *v
	}
	if v := item.SentAt; v != nil {
		offer["sent_at"] = *v
	}
	if v := item.StartsAt; v != nil {
		offer["starts_at"] = *v
	}
	if v := item.Status; v != nil {
		offer["status"] = *v
	}
	if v := item.Version; v != nil {
		offer["version"] = *v
	}
	return offer
}

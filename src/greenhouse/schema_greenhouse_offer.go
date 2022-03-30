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
			Optional: true,
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
		"id": {
			Type:     schema.TypeInt,
			Optional: true,
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
	offer["application_id"] = item.ApplicationId
	offer["candidate_id"] = item.CandidateId
	offer["created_at"] = item.CreatedAt
	offer["custom_fields"] = item.CustomFields
	offer["job_id"] = item.JobId
	offer["keyed_custom_fields"] = item.KeyedCustomFields
	offer["opening"] = flattenJobOpening(ctx, item.Opening)
	offer["resolved_at"] = item.ResolvedAt
	offer["sent_at"] = item.SentAt
	offer["starts_at"] = item.StartsAt
	offer["status"] = item.Status
	offer["version"] = item.Version
	return offer
}

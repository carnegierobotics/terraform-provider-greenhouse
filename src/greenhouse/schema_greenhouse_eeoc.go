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
	answer["description"] = item.Description
	answer["id"] = strconv.Itoa(item.Id)
	answer["message"] = item.Message
	return answer
}

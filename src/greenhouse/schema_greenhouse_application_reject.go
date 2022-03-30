package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApplicationReject() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"notes": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"rejection_email": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseRejectionEmail(),
			},
		},
		"rejection_reason": {
			Type:     schema.TypeInt,
			Computed: true,
		},
	}
}

func inflateApplicationRejects(ctx context.Context, source *[]interface{}) (*[]greenhouse.ApplicationReject, diag.Diagnostics) {
	list := make([]greenhouse.ApplicationReject, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApplicationReject(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApplicationReject(ctx context.Context, source *map[string]interface{}) (*greenhouse.ApplicationReject, diag.Diagnostics) {
	var obj greenhouse.ApplicationReject
	if v, ok := (*source)["notes"].(string); ok && len(v) > 0 {
		obj.Notes = v
	}
	if v, ok := (*source)["rejection_email"].([]interface{}); ok && len(v) > 0 {
		item, diagErr := inflateRejectionEmails(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.RejectionEmail = &(*item)[0]
	}
	if v, ok := (*source)["rejection_reason"].(int); ok {
		obj.RejectionReasonId = v
	}
	return &obj, nil
}

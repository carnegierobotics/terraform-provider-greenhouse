package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApplicationHire() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"close_reason_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"opening_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateApplicationHires(ctx context.Context, source *[]interface{}) (*[]greenhouse.ApplicationHire, diag.Diagnostics) {
	list := make([]greenhouse.ApplicationHire, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateApplicationHire(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateApplicationHire(ctx context.Context, source *map[string]interface{}) (*greenhouse.ApplicationHire, diag.Diagnostics) {
	var obj greenhouse.ApplicationHire
	if v, ok := (*source)["close_reason_id"].(int); ok {
		obj.CloseReasonId = &v
	}
	if v, ok := (*source)["opening_id"].(int); ok {
		obj.OpeningId = &v
	}
	if v, ok := (*source)["start_date"].(string); ok && len(v) > 0 {
		obj.StartDate = &v
	}
	return &obj, nil
}

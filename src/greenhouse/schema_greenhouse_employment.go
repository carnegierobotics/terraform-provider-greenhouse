package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmployment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"company_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"end_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"start_date": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"title": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateEmployments(ctx context.Context, source *[]interface{}) (*[]greenhouse.Employment, diag.Diagnostics) {
	list := make([]greenhouse.Employment, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		item, diagErr := inflateEmployment(ctx, &itemMap)
		if diagErr != nil {
			return nil, diagErr
		}
		list[i] = *item
	}
	return &list, nil
}

func inflateEmployment(ctx context.Context, source *map[string]interface{}) (*greenhouse.Employment, diag.Diagnostics) {
	var obj greenhouse.Employment
	if v, ok := (*source)["company_name"].(string); ok && len(v) > 0 {
		obj.CompanyName = &v
	}
	if v, ok := (*source)["end_date"].(string); ok && len(v) > 0 {
		obj.EndDate = &v
	}
	if v, ok := (*source)["start_date"].(string); ok && len(v) > 0 {
		obj.StartDate = &v
	}
	if v, ok := (*source)["title"].(string); ok && len(v) > 0 {
		obj.Title = &v
	}
	return &obj, nil
}

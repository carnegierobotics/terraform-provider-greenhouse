package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseActivityFeed() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"activities": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseActivity(),
			},
		},
		"candidate_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"emails": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseEmail(),
			},
		},
		"notes": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseNote(),
			},
		},
	}
}

func inflateActivityFeeds(ctx context.Context, source *[]interface{}) (*[]greenhouse.ActivityFeed, diag.Diagnostics) {
	list := make([]greenhouse.ActivityFeed, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateActivityFeed(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateActivityFeed(ctx context.Context, source *map[string]interface{}) (*greenhouse.ActivityFeed, diag.Diagnostics) {
	var obj greenhouse.ActivityFeed
	if v, ok := (*source)["activities"].([]interface{}); ok && len(v) > 0 {
		activities, diagErr := inflateActivities(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Activities = *activities
	}
	if v, ok := (*source)["emails"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateEmails(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Emails = *list
	}
	if v, ok := (*source)["notes"].([]interface{}); ok && len(v) > 0 {
		list, diagErr := inflateNotes(ctx, &v)
		if diagErr != nil {
			return nil, diagErr
		}
		obj.Notes = *list
	}
	return &obj, nil
}

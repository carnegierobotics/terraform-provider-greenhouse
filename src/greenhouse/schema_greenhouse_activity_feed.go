package greenhouse

import (
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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

func convertToActivityFeedList(list []interface{}) *[]greenhouse.ActivityFeed {
	newList := make([]greenhouse.ActivityFeed, len(list))
	for i := range list {
		newList[i] = list[i].(greenhouse.ActivityFeed)
	}
	return &newList
}

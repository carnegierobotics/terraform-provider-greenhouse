package greenhouse

import (
  "github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOffice() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"location": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"primary_contact_user_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"child_ids": {
			Type:     schema.TypeSet,
      Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	}
}

func flattenOffices(list *[]greenhouse.Office) []interface{} {
  if list != nil {
    flatList := make([]interface{}, len(*list), len(*list))
    for i, item := range *list {
      office := make(map[string]interface{})
      office["name"] = item.Name
      office["location"] = flattenLocation(&item.Location)
      office["primary_context_user_id"] = item.PrimaryContactUserId
      office["parent_id"] = item.ParentId
      office["child_ids"] = item.ChildIds
      flatList[i] = office
    }
    return flatList
  }
  return make([]interface{}, 0)
}

func flattenLocation(item *greenhouse.Location) map[string]interface{} {
  location := make(map[string]interface{})
  if item.Name != "" {
    location["name"] = item.Name
  }
  return location
}

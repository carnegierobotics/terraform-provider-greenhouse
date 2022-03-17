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
		"location_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"location": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseLocation(),
			},
		},
		"primary_contact_user_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"parent_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"child_ids": {
			Type:     schema.TypeSet,
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
			office["primary_contact_user_id"] = item.PrimaryContactUserId
			office["parent_id"] = item.ParentId
			office["child_ids"] = item.ChildIds
			flatList[i] = office
		}
		return flatList
	}
	return make([]interface{}, 0)
}

package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOffice() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"child_ids": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"location": {
			/* This is how it should be, but state is not ready.
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseLocation(),
			},
			*/
			//So this is how it is...
			Type:     schema.TypeMap,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"location_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
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
	}
}

func flattenOffices(ctx context.Context, list *[]greenhouse.Office) []interface{} {
	tflog.Debug(ctx, "Flattening office list", "officelist", fmt.Sprintf("%+v", list))
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenOffice(ctx, &item)
		}
		tflog.Debug(ctx, "Flattened office list", "officelist", fmt.Sprintf("%+v", flatList))
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenOffice(ctx context.Context, item *greenhouse.Office) map[string]interface{} {
	tflog.Debug(ctx, "Flattening office", "office", fmt.Sprintf("%+v", item))
	office := make(map[string]interface{})
	office["name"] = item.Name
	office["location"] = flattenLocation(ctx, item.Location)[0]
	office["primary_contact_user_id"] = item.PrimaryContactUserId
	office["parent_id"] = item.ParentId
	office["child_ids"] = item.ChildIds
	tflog.Debug(ctx, "Flattened office", "office", fmt.Sprintf("%+v", office))
	return office
}

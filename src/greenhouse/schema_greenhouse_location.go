package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseLocation() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func inflateLocations(ctx context.Context, source *[]interface{}) (*[]greenhouse.Location, diag.Diagnostics) {
	list := make([]greenhouse.Location, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateLocation(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateLocation(ctx context.Context, source *map[string]interface{}) (*greenhouse.Location, diag.Diagnostics) {
	var obj greenhouse.Location
	/*
	  if v, ok := (*source)["address"].(string); ok && len(v) > 0 {
	    obj.Address = v
	  }
	*/
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = &v
	}
	return &obj, nil
}

func flattenLocation(ctx context.Context, item *greenhouse.Location) []interface{} {
	tflog.Trace(ctx, "Flattening location", "location", fmt.Sprintf("%+v", item))
	location := make([]interface{}, 1, 1)
	oneLocation := make(map[string]interface{})
	if v := item.Name; v != nil {
		oneLocation["name"] = *v
	}
	location[0] = oneLocation
	tflog.Trace(ctx, "Flattened location", "location", fmt.Sprintf("%+v", location))
	return location
}

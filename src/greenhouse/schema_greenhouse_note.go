package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseNote() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"body": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"user": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
		"visiblity": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"visibility": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateNotes(ctx context.Context, source *[]interface{}) (*[]greenhouse.Note, diag.Diagnostics) {
	list := make([]greenhouse.Note, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateNote(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateNote(ctx context.Context, source *map[string]interface{}) (*greenhouse.Note, diag.Diagnostics) {
	var obj greenhouse.Note
	if v, ok := (*source)["body"].(string); ok && len(v) > 0 {
		obj.Body = &v
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
	}
	if v, ok := (*source)["private"].(bool); ok {
		obj.Private = &v
	}
	if v, ok := (*source)["user"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateUsers(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.User = &(*list)[0]
	}
	if v, ok := (*source)["user_id"].(int); ok {
		obj.UserId = &v
	}
	if v, ok := (*source)["visibility"].(string); ok && len(v) > 0 {
		obj.Visibility = &v
	}
	return &obj, nil
}

func flattenNotes(ctx context.Context, list *[]greenhouse.Note) []interface{} {
	if list != nil {
		tflog.Trace(ctx, "Flattening notes.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, note := range *list {
			note := flattenNote(ctx, &note)
			flatList[i] = note
		}
		tflog.Trace(ctx, "Finished flattening notes.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenNote(ctx context.Context, item *greenhouse.Note) map[string]interface{} {
	tflog.Trace(ctx, "Flattening one note.")
	note := make(map[string]interface{})
	if v := item.Body; v != nil {
		note["body"] = *v
	}
	if v := item.CreatedAt; v != nil {
		note["created_at"] = *v
	}
	if v := item.Private; v != nil {
		note["private"] = *v
	}
	if v := item.User; v != nil {
		note["user"] = flattenUsersBasics(ctx, &[]greenhouse.User{*v})
	}
	if v := item.Visibility; v != nil {
		note["visibility"] = *v
	}
	tflog.Trace(ctx, "Finished flattening note.")
	return note
}

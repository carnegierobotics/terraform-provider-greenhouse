package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseNote() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"body": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"user": {
			Type:     schema.TypeSet,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseUserBasics(),
			},
		},
		"private": {
			Type:     schema.TypeBool,
			Computed: true,
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

func flattenNotes(ctx context.Context, list *[]greenhouse.Note) []interface{} {
	if list != nil {
		tflog.Debug(ctx, "Flattening notes.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, note := range *list {
			note := flattenNote(ctx, &note)
			flatList[i] = note
		}
		tflog.Debug(ctx, "Finished flattening notes.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenNote(ctx context.Context, item *greenhouse.Note) map[string]interface{} {
	tflog.Debug(ctx, "Flattening one note.")
	note := make(map[string]interface{})
	note["body"] = item.Body
	note["created_at"] = item.CreatedAt
	note["private"] = item.Private
	note["user"] = flattenUsersBasics(ctx, &[]greenhouse.User{*item.User})
	note["visibility"] = item.Visibility
	tflog.Debug(ctx, "Finished flattening note.")
	return note
}

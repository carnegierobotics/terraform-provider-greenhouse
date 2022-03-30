package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"body": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cc": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"from": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"to": {
			Type:     schema.TypeString,
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
	}
}

func inflateEmails(ctx context.Context, source *[]interface{}) (*[]greenhouse.Email, diag.Diagnostics) {
	list := make([]greenhouse.Email, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateEmail(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateEmail(ctx context.Context, source *map[string]interface{}) (*greenhouse.Email, diag.Diagnostics) {
	var obj greenhouse.Email
	if v, ok := (*source)["body"].(string); ok && len(v) > 0 {
		obj.Body = v
	}
	if v, ok := (*source)["cc"].(string); ok && len(v) > 0 {
		obj.Cc = v
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = v
	}
	if v, ok := (*source)["from"].(string); ok && len(v) > 0 {
		obj.From = v
	}
	if v, ok := (*source)["subject"].(string); ok && len(v) > 0 {
		obj.Subject = v
	}
	if v, ok := (*source)["to"].(string); ok && len(v) > 0 {
		obj.To = v
	}
	if v, ok := (*source)["user"].([]interface{}); ok && len(v) > 0 {
		item, err := inflateUser(ctx, &(v[0]))
		if err != nil {
			return nil, err
		}
		obj.User = item
	}
	return &obj, nil
}

func flattenEmails(ctx context.Context, list *[]greenhouse.Email) []interface{} {
	if list != nil {
		tflog.Debug(ctx, "Flattening emails.")
		flatList := make([]interface{}, len(*list), len(*list))
		for i, email := range *list {
			email := flattenEmail(ctx, &email)
			flatList[i] = email
		}
		tflog.Debug(ctx, "Finished flattening emails.")
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenEmail(ctx context.Context, item *greenhouse.Email) map[string]interface{} {
	tflog.Debug(ctx, "Flattening one email.")
	email := make(map[string]interface{})
	email["body"] = item.Body
	email["cc"] = item.Cc
	email["created_at"] = item.CreatedAt
	email["subject"] = item.Subject
	email["to"] = item.To
	email["user"] = flattenUsersBasics(ctx, &[]greenhouse.User{*item.User})
	tflog.Debug(ctx, "Finished flattening email.")
	return email
}

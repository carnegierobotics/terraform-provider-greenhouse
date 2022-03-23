package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseEmail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subject": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"body": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"to": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"from": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"cc": {
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
	}
}

func flattenEmails(ctx context.Context, list *[]greenhouse.Email) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, email := range *list {
			email := flattenEmail(ctx, &email)
			flatList[i] = email
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenEmail(ctx context.Context, item *greenhouse.Email) map[string]interface{} {
	email := make(map[string]interface{})
	email["body"] = item.Body
	email["cc"] = item.Cc
	email["created_at"] = item.CreatedAt
	email["subject"] = item.Subject
	email["to"] = item.To
	email["user"] = flattenUserBasics(ctx, item.User)
	return email
}

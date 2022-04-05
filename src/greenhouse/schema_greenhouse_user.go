package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseUser() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"disable_user": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"disabled": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"emails": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"employee_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"first_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"last_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"linked_candidate_ids": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"primary_email_address": {
			Type:     schema.TypeString,
			Required: true,
		},
		"send_email": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"site_admin": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateUsers(ctx context.Context, source *[]interface{}) (*[]greenhouse.User, diag.Diagnostics) {
	tflog.Debug(ctx, fmt.Sprintf("Inflating users: %+v", source))
	if source != nil && len(*source) > 0 {
		var list []greenhouse.User
		err := convertType(ctx, *source, list)
		if err != nil {
			return nil, err
		}
		return &list, nil
	}
	return nil, nil
}

func inflateUser(ctx context.Context, source interface{}) (*greenhouse.User, diag.Diagnostics) {
	var item greenhouse.User
	err := convertType(ctx, source, item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func flattenUser(ctx context.Context, item *greenhouse.User) map[string]interface{} {
	tflog.Debug(ctx, "User item:", fmt.Sprintf("%+v\n", *item))
	user := make(map[string]interface{})
	if v := item.CreatedAt; v != nil {
		user["created_at"] = *v
	}
	if v := item.Disabled; v != nil {
		user["disabled"] = *v
	}
	if v := item.Emails; len(v) > 0 {
		user["emails"] = item.Emails
	}
	if v := item.EmployeeId; v != nil {
		user["employee_id"] = *v
	}
	if v := item.FirstName; v != nil {
		user["first_name"] = *v
	}
	if v := item.LastName; v != nil {
		user["last_name"] = *v
	}
	if v := item.LinkedCandidateIds; len(v) > 0 {
		user["linked_candidate_ids"] = item.LinkedCandidateIds
	}
	if v := item.Name; v != nil {
		user["name"] = item.Name
	}
	if v := item.PrimaryEmail; v != nil {
		user["primary_email_address"] = item.PrimaryEmail
	}
	if v := item.SiteAdmin; v != nil {
		user["site_admin"] = item.SiteAdmin
	}
	if v := item.UpdatedAt; v != nil {
		user["updated_at"] = item.UpdatedAt
	}
	return user
}

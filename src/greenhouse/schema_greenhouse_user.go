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
	user := make(map[string]interface{})
	user["created_at"] = item.CreatedAt
	user["disabled"] = item.Disabled
	user["emails"] = item.Emails
	user["employee_id"] = item.EmployeeId
	user["first_name"] = item.FirstName
	user["last_name"] = item.LastName
	user["linked_candidate_ids"] = item.LinkedCandidateIds
	user["name"] = item.Name
	user["primary_email_address"] = item.PrimaryEmail
	user["site_admin"] = item.SiteAdmin
	user["updated_at"] = item.UpdatedAt
	return user
}

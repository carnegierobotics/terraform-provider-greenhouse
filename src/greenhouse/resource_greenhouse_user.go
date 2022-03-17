package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseUserCreate,
		ReadContext:   resourceGreenhouseUserRead,
		UpdateContext: resourceGreenhouseUserUpdate,
		DeleteContext: resourceGreenhouseUserDelete,
		Exists:        resourceGreenhouseUserExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseUser(),
	}
}

func resourceGreenhouseUserExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/users/%d", id))
}

func resourceGreenhouseUserCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.UserCreateInfo{
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Email:     d.Get("primary_email_address").(string),
		SendEmail: d.Get("send_email").(bool),
	}
	id, err := greenhouse.CreateUser(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseUserUpdate(ctx, d, meta)
}

func resourceGreenhouseUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetUser(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.Set("name", obj.Name)
	d.Set("first_name", obj.FirstName)
	d.Set("last_name", obj.LastName)
	d.Set("employee_id", obj.EmployeeId)
	d.Set("primary_email_address", obj.PrimaryEmail)
	d.Set("updated_at", obj.UpdatedAt)
	d.Set("created_at", obj.CreatedAt)
	d.Set("disabled", obj.Disabled)
	d.Set("site_admin", obj.SiteAdmin)
	d.Set("emails", obj.Emails)
	d.Set("linked_candidate_ids", obj.LinkedCandidateIds)
	return nil
}

func resourceGreenhouseUserUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if d.HasChange("disable_user") {
		if d.Get("disable_user").(bool) {
			err = greenhouse.DisableUser(meta.(*greenhouse.Client), ctx, id)
		} else {
			err = greenhouse.EnableUser(meta.(*greenhouse.Client), ctx, id)
		}
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	} else {
		updateObject := greenhouse.UserUpdateInfo{
			FirstName: d.Get("first_name").(string),
			LastName:  d.Get("last_name").(string),
		}
		err = greenhouse.UpdateUser(meta.(*greenhouse.Client), ctx, id, &updateObject)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	return resourceGreenhouseUserRead(ctx, d, meta)
}

func resourceGreenhouseUserDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for users."}}
}

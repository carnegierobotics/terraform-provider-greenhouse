package greenhouse

import (
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceGreenhouseUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreenhouseUserCreate,
		Read:   resourceGreenhouseUserRead,
		Update: resourceGreenhouseUserUpdate,
		Delete: resourceGreenhouseUserDelete,
		Exists: resourceGreenhouseUserExists,
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
	return greenhouse.Exists(meta.(*greenhouse.Client), "users", id)
}

func resourceGreenhouseUserCreate(d *schema.ResourceData, meta interface{}) error {
	createObject := greenhouse.UserCreateInfo{
		FirstName: d.Get("first_name").(string),
		LastName:  d.Get("last_name").(string),
		Email:     d.Get("primary_email_address").(string),
		SendEmail: d.Get("send_email").(bool),
	}
	id, err := greenhouse.CreateUser(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return err
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseUserRead(d, meta)
}

func resourceGreenhouseUserRead(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	obj, err := greenhouse.GetUser(meta.(*greenhouse.Client), id)
	if err != nil {
		return err
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

func resourceGreenhouseUserUpdate(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	if d.HasChange("disable_user") {
		if d.Get("disable_user").(bool) {
			err = greenhouse.DisableUser(meta.(*greenhouse.Client), id)
		} else {
			err = greenhouse.EnableUser(meta.(*greenhouse.Client), id)
		}
		if err != nil {
			return err
		}
	} else {
		updateObject := greenhouse.UserUpdateInfo{
			FirstName: d.Get("first_name").(string),
			LastName:  d.Get("last_name").(string),
		}
		err = greenhouse.UpdateUser(meta.(*greenhouse.Client), id, &updateObject)
		if err != nil {
			return err
		}
	}
	return resourceGreenhouseUserRead(d, meta)
}

func resourceGreenhouseUserDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: delete is not supported for users.")
}

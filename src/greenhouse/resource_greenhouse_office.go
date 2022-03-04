package greenhouse

import (
  "context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceGreenhouseOffice() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreenhouseOfficeCreate,
		Read:   resourceGreenhouseOfficeRead,
		Update: resourceGreenhouseOfficeUpdate,
		Delete: resourceGreenhouseOfficeDelete,
		Exists: resourceGreenhouseOfficeExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseOffice(),
	}
}

func resourceGreenhouseOfficeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "offices", id, context.TODO())
}

func resourceGreenhouseOfficeCreate(d *schema.ResourceData, meta interface{}) error {
	createObject := greenhouse.OfficeCreateInfo{
		Name:                 d.Get("name").(string),
		Location:             d.Get("location.name").(string),
		PrimaryContactUserId: d.Get("primary_contact_user_id").(int),
		ParentId:             d.Get("parent_id").(int),
	}
	id, err := greenhouse.CreateOffice(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return err
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseOfficeRead(d, meta)
}

func resourceGreenhouseOfficeRead(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	obj, err := greenhouse.GetOffice(meta.(*greenhouse.Client), id)
	if err != nil {
		return err
	}
	d.Set("name", obj.Name)
	d.Set("location", obj.Location)
	d.Set("primary_contact_user_id", obj.PrimaryContactUserId)
	d.Set("parent_id", obj.ParentId)
	d.Set("child_ids", obj.ChildIds)
	return nil
}

func resourceGreenhouseOfficeUpdate(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	updateObject := greenhouse.OfficeUpdateInfo{
		Name:     d.Get("name").(string),
		Location: d.Get("location.name").(string),
	}
	err = greenhouse.UpdateOffice(meta.(*greenhouse.Client), id, &updateObject)
	if err != nil {
		return err
	}
	return resourceGreenhouseOfficeRead(d, meta)
}

func resourceGreenhouseOfficeDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: delete is not supported for offices.")
}

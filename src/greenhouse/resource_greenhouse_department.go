package greenhouse

import (
  "context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

type ReadFunc func(d *schema.ResourceData, m interface{}) error

func resourceGreenhouseDepartment() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreenhouseDepartmentCreate,
		Read:   resourceGreenhouseDepartmentRead,
		Update: resourceGreenhouseDepartmentUpdate,
		Delete: resourceGreenhouseDepartmentDelete,
		Exists: resourceGreenhouseDepartmentExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseDepartment(),
	}
}

func resourceGreenhouseDepartmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "departments", id, context.TODO())
}

func resourceGreenhouseDepartmentCreate(d *schema.ResourceData, meta interface{}) error {
	createObject := greenhouse.DepartmentCreateInfo{
		Name:     d.Get("name").(string),
		ParentId: d.Get("parent_id").(int),
	}
	id, err := greenhouse.CreateDepartment(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return err
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseDepartmentRead(d, meta)
}

func resourceGreenhouseDepartmentRead(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	obj, err := greenhouse.GetDepartment(meta.(*greenhouse.Client), id)
	if err != nil {
		return err
	}
	d.Set("name", obj.Name)
	d.Set("parent_id", obj.ParentId)
	d.Set("child_ids", obj.ChildIds)
	return nil
}

func resourceGreenhouseDepartmentUpdate(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	updateObject := greenhouse.DepartmentUpdateInfo{
		Name: d.Get("name").(string),
	}
	err = greenhouse.UpdateDepartment(meta.(*greenhouse.Client), id, &updateObject)
	if err != nil {
		return err
	}
	return resourceGreenhouseDepartmentRead(d, meta)
}

func resourceGreenhouseDepartmentDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: delete is not supported for departments.")
}

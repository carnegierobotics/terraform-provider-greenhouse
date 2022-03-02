package greenhouse

import (
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceGreenhouseCloseReason() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreenhouseCloseReasonCreate,
		Read:   resourceGreenhouseCloseReasonRead,
		Update: resourceGreenhouseCloseReasonUpdate,
		Delete: resourceGreenhouseCloseReasonDelete,
		Exists: resourceGreenhouseCloseReasonExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseCloseReason(),
	}
}

func resourceGreenhouseCloseReasonExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "close_reasons", id)
}

func resourceGreenhouseCloseReasonCreate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: create is not supported for close_reasons.")
}

func resourceGreenhouseCloseReasonRead(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	obj, err := greenhouse.GetCloseReason(meta.(*greenhouse.Client), id)
	if err != nil {
		return err
	}
	d.Set("name", obj.Name)
	return nil
}

func resourceGreenhouseCloseReasonUpdate(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: update is not supported for close_reasons.")
}

func resourceGreenhouseCloseReasonDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: delete is not supported for close_reasons.")
}

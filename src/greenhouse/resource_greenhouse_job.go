package greenhouse

import (
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceGreenhouseJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceGreenhouseJobCreate,
		Read:   resourceGreenhouseJobRead,
		Update: resourceGreenhouseJobUpdate,
		Delete: resourceGreenhouseJobDelete,
		Exists: resourceGreenhouseJobExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseJob(),
	}
}

func resourceGreenhouseJobExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "jobs", id)
}

func resourceGreenhouseJobCreate(d *schema.ResourceData, meta interface{}) error {
	createObject := greenhouse.JobCreateInfo{
		TemplateJobId:  d.Get("template_job_id").(int),
		NumberOpenings: d.Get("number_of_openings").(int),
		JobPostName:    d.Get("job_post_name").(string),
		JobName:        d.Get("job_name").(string),
		DepartmentId:   d.Get("department_id").(int),
		OfficeIds:      convertListIToListD(d.Get("office_ids").(*schema.Set).List()),
		RequisitionId:  d.Get("requisition_id").(string),
		OpeningIds:     convertListIToListA(d.Get("opening_ids").(*schema.Set).List()),
	}
	id, err := greenhouse.CreateJob(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return err
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseJobRead(d, meta)
}

func convertListIToListD(list []interface{}) []int {
  var newList []int
  for i := range list {
    newList[i] = list[i].(int)
  }
  return newList
}

func convertListIToListA(list []interface{}) []string {
  var newList []string
  for i := range list {
    newList[i] = list[i].(string)
  }
  return newList
}

func resourceGreenhouseJobRead(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	obj, err := greenhouse.GetJob(meta.(*greenhouse.Client), id)
	if err != nil {
		return err
	}
	d.Set("job_name", obj.Name)
	d.Set("departments", obj.Departments)
	d.Set("offices", obj.Offices)
	d.Set("requisition_id", obj.RequisitionId)
	d.Set("openings", obj.Openings)
	return nil
}

func resourceGreenhouseJobUpdate(d *schema.ResourceData, meta interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}
	updateObject := greenhouse.JobUpdateInfo{
		Name:                     d.Get("name").(string),
		Notes:                    d.Get("notes").(string),
		Anywhere:                 d.Get("anywhere").(bool),
		RequisitionId:            d.Get("requisition_id").(string),
		TeamsandResponsibilities: d.Get("teams_and_responsibilities").(string),
		HowToSellThisJob:         d.Get("how_to_sell_this_job").(string),
		OfficeIds:                d.Get("office_ids").([]int),
		DepartmentId:             d.Get("department_id").(int),
	}
	err = greenhouse.UpdateJob(meta.(*greenhouse.Client), id, &updateObject)
	if err != nil {
		return err
	}
	return resourceGreenhouseJobRead(d, meta)
}

func resourceGreenhouseJobDelete(d *schema.ResourceData, meta interface{}) error {
	return fmt.Errorf("Error: delete is not supported for jobs.")
}

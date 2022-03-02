package greenhouse

import (
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func dataSourceGreenhouseDepartment() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGreenhouseDepartmentRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"child_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			/* Not in our product tier
			   "parent_department_external_id": {
			     Type: schema.TypeString,
			     Optional: true,
			     Computed: true,
			   },
			   "child_department_external_ids": {
			     Type: schema.TypeSet,
			     Optional: true,
			     Computed: true,
			     Elem: &schema.Schema {
			       Type: schema.TypeString,
			     }
			   },
			   "external_id": {
			     Type: schema.TypeString,
			     Optional: true,
			   }
			*/
		},
	}
}

func dataSourceGreenhouseDepartmentRead(d *schema.ResourceData, meta interface{}) error {
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

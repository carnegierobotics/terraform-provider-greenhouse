package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func schemaGreenhouseDepartment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
			Computed: true,
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
	}
}

package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseProspectApplication() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "job_ids": {
      Type: schema.TypeSet,
      Optional: true,
      Elem: &schema.Schema{
        Type: schema.TypeInt,
      },
    },
    "prospect": {
      Type: schema.TypeBool,
      Optional: true,
    },
    "prospective_department_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "prospective_office_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "prospect_owner_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "prospect_pool_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "prospect_pool_stage_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "referrer": {
      Type: schema.TypeSet,
      MaxItems: 1,
      Optional: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseTypeTypeValue(),
      },
    },
    "source_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
	}
}

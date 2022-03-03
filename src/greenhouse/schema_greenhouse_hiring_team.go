package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func schemaGreenhouseHiringTeam() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
    "team": {
      Type:     schema.TypeMap,
      Required: true,
      Elem:     &schema.Schema{
        Type: schema.TypeSet,
        Required: true,
        Elem: &schema.Resource{
            Schema: schemaGreenhouseHiringMember(),
        },
      },
    },
	}
}

func schemaGreenhouseHiringMember() map[string]*schema.Schema {
  return map[string]*schema.Schema{
    "user_id": {
      Type:     schema.TypeString,
      Required: true,
    },
    "active": {
      Type:     schema.TypeBool,
      Computed: true,
    },
    "responsible": {
      Type:     schema.TypeBool,
      Computed: true,
    },
    "responsible_for_future_candidates": {
      Type: schema.TypeBool,
      Optional: true,
    },
    "responsible_for_active_candidates": {
      Type: schema.TypeBool,
      Optional: true,
    },
    "responsible_for_inactive_candidates": {
      Type: schema.TypeBool,
      Optional: true,
    },
    "first_name": {
      Type: schema.TypeString,
      Computed: true,
    },
    "last_name": {
      Type: schema.TypeString,
      Computed: true,
    },
    "name": {
      Type:     schema.TypeString,
      Computed: true,
    },
    "employee_id": {
      Type:     schema.TypeString,
      Computed: true,
    },
  }
}

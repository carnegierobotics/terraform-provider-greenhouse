package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseInterview() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "default_interviewer_users": {
      Type: schema.TypeSet,
      Optional: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseInterviewer(),
      },
    },
    "estimated_minutes": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "interview_kit": {
      Type: schema.TypeSet,
      MaxItems: 1,
      Optional: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseInterviewKit(),
      },
    },
    "name": {
      Type: schema.TypeString,
      Optional: true,
    },
    "schedulable": {
      Type: schema.TypeBool,
      Optional: true,
    },
	}
}

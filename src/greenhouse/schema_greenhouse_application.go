package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseApplication() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"candidate_id": {
			Type:     schema.TypeInt,
      Description: "The ID of the candidate applying for this job.",
			Computed: true,
		},
    "prospect": {
      Type:     schema.TypeBool,
      Description: "The candidate is a prospect and has not yet applied.",
      Computed: true,
    },
    "applied_at": {
      Type: schema.TypeString,
      Description: "The date of the application.",
      Computed: true,
    },
    "rejected_at": {
      Type: schema.TypeString,
      Description: "The date of the application's rejection.",
      Computed: true,
    },
    "last_activity_at": {
      Type: schema.TypeString,
      Description: "The date of the application's last activity.",
      Computed: true,
    },
    "location": {
      Type: schema.TypeList,
      Description: "The contents of a location question on a job post.",
      MaxItems: 1,
      Computed: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseLocation(),
      },
    },
    "credited_to": {
      Type:     schema.TypeSet,
      Description: "The user who will receive credit for this application.",
      MaxItems: 1,
      Computed: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseUserBasics(),
      },
    },
    "private": {
      Type: schema.TypeBool,
      Computed: true,
    },
    "visiblity": {
      Type: schema.TypeString,
      Computed: true,
    },
    "visibility": {
      Type: schema.TypeString,
      Computed: true,
    },
    "status": {
      Type: schema.TypeString,
      Computed: true,
    },
	}
}

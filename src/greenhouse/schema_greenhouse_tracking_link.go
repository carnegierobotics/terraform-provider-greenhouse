package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseTrackingLink() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "created_at": {
      Type: schema.TypeString,
      Computed: true,
    },
    "credited_to": {
      Type: schema.TypeSet,
      MaxItems: 1,
      Optional: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseUser(),
      },
    },
    "job_board": {
      Type: schema.TypeSet,
      MaxItems: 1,
      Optional: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseJobBoard(),
      },
    },
    "job_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "job_post_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "related_post_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "related_post_type": {
      Type: schema.TypeString,
      Optional: true,
    },
    "source": {
      Type: schema.TypeSet,
      MaxItems: 1,
      Optional: true,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseSource(),
      },
    },
    "token": {
      Type: schema.TypeString,
      Optional: true,
    },
    "updated_at": {
      Type: schema.TypeString,
      Computed: true,
    },
	}
}

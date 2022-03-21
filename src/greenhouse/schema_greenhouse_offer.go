package greenhouse

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseOffer() map[string]*schema.Schema {
	return map[string]*schema.Schema{
    "application_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "candidate_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "created_at": {
      Type: schema.TypeString,
      Computed: true,
    },
    "custom_fields": {
      Type: schema.TypeMap,
      Optional: true,
      Elem: &schema.Schema{
        Type: schema.TypeString,
      },
    },
    "job_id": {
      Type: schema.TypeInt,
      Optional: true,
    },
    "keyed_custom_fields": {
      Type: schema.TypeSet,
      Optional: true,
      MaxItems: 1,
      Elem: &schema.Schema{
        Type: schema.TypeMap,
        Elem: &schema.Resource{
          Schema: schemaGreenhouseKeyedCustomField(),
        },
      },
    },
    "opening": {
      Type: schema.TypeSet,
      Optional: true,
      MaxItems: 1,
      Elem: &schema.Resource{
        Schema: schemaGreenhouseJobOpening(),
      },
    },
    "resolved_at": {
      Type: schema.TypeString,
      Computed: true,
    },
    "sent_at": {
      Type: schema.TypeString,
      Computed: true,
    },
    "starts_at": {
      Type: schema.TypeString,
      Computed: true,
    },
    "status": {
      Type: schema.TypeString,
      Optional: true,
    },
    "version": {
      Type: schema.TypeInt,
      Computed: true,
    },
	}
}

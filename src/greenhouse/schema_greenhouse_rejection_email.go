package greenhouse

import (
  "context"
  "github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseRejectionEmail() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"email_template_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"send_email_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateRejectionEmails(ctx context.Context, source *[]interface{}) (*[]greenhouse.RejectionEmail, diag.Diagnostics) {
  list := make([]greenhouse.RejectionEmail, len(*source), len(*source))
  for i, item := range *source {
    itemMap := item.(map[string]interface{})
    obj, err := inflateRejectionEmail(ctx, &itemMap)
    if err != nil {
      return nil, err
    }
    list[i] = *obj
  }
  return &list, nil
}

func inflateRejectionEmail(ctx context.Context, source *map[string]interface{}) (*greenhouse.RejectionEmail, diag.Diagnostics) {
  var obj greenhouse.RejectionEmail
  if v, ok := (*source)["email_template_id"].(string); ok && len(v) > 0 {
    obj.EmailTemplateId = v
  }
  if v, ok := (*source)["send_email_at"].(string); ok && len(v) > 0 {
    obj.SendEmailAt = v
  }
  return &obj, nil
}

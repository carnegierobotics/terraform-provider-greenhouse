package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseEmailTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseEmailTemplateRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseEmailTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  name := d.Get("name").(string)
	list, err := greenhouse.GetAllEmailTemplates(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for _, email := range *list {
		if email.Name == name {
			d.SetId(strconv.Itoa(email.Id))
      d.Set("body", email.Body)
      d.Set("cc", email.Cc)
      d.Set("created_at", email.CreatedAt)
      d.Set("default", email.Default)
      d.Set("description", email.Description)
      d.Set("from", email.From)
      d.Set("html_body", email.HtmlBody)
      d.Set("type", email.Type)
      d.Set("updated_at", email.UpdatedAt)
      d.Set("user", flattenUser(ctx, email.User))
			return nil
		}
	}
	return nil
}

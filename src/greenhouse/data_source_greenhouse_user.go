package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "strconv"
)

func dataSourceGreenhouseUser() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseUserRead,
		Schema: map[string]*schema.Schema{
      "name": {
        Type: schema.TypeString,
        Required: true,
      },
    },
	}
}

func dataSourceGreenhouseUserRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllUsers(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  name := d.Get("name").(string)
  for _, user := range *list {
    if user.Name == name {
      d.SetId(strconv.Itoa(user.Id))
      return nil
    }
  }
	return nil
}

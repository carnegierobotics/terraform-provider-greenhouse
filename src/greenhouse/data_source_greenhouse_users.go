package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseUsersRead,
		Schema: map[string]*schema.Schema{
      "names": {
        Type: schema.TypeList,
        Computed: true,
        Elem: &schema.Schema{
          Type: schema.TypeString,
        },
      },
    },
	}
}

func dataSourceGreenhouseUsersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllUsers(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  users := make([]string, len(*list), len(*list))
  for i, user := range *list {
    users[i] = user.Name
  }
  d.SetId("all")
	d.Set("names", users)
	return nil
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "strconv"
)

func dataSourceGreenhouseUserRole() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceGreenhouseUserRoleRead,
		Schema: schemaGreenhouseUserRole(),
	}
}

func dataSourceGreenhouseUserRoleRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  name := d.Get("name").(string)
	list, err := greenhouse.GetAllUserRoles(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  for _, role := range *list {
    if role.Name == name {
      d.SetId(strconf.Itoa(role.Id))
      return nil
    }
  }
	return nil
}

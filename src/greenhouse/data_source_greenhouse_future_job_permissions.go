package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseFutureJobPermissions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseFutureJobPermissionsRead,
		Schema: map[string]*schema.Schema{
			"permissions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseFutureJobPermission(),
				},
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseFutureJobPermissionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetFutureJobPermissions(meta.(*greenhouse.Client), ctx, d.Get("user_id").(int))
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("all")
	d.Set("permissions", flattenFutureJobPermissions(ctx, list))
	return nil
}

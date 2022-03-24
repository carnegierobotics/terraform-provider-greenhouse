package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseJobPermissions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseJobPermissionsRead,
		Schema: map[string]*schema.Schema{
			"permissions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseUserPermission(),
				},
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseJobPermissionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetJobPermissions(meta.(*greenhouse.Client), ctx, d.Get("user_id").(int))
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("all")
	d.Set("permissions", flattenUserPermissions(ctx, list))
	return nil
}

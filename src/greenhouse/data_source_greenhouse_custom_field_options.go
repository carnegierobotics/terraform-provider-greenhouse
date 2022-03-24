package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseCustomFieldOptions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCustomFieldOptionsRead,
		Schema: map[string]*schema.Schema{
			"custom_field_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"options": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseCustomFieldOption(),
				},
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGreenhouseCustomFieldOptionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("custom_field_id").(int)
	typeStr := d.Get("type").(string)
	obj, err := greenhouse.GetCustomFieldOptions(meta.(*greenhouse.Client), ctx, id, typeStr)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	d.Set("options", flattenCustomFieldOptions(ctx, obj))
	return nil
}

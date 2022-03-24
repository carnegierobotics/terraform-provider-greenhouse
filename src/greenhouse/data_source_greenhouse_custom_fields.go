package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseCustomFields() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCustomFieldsRead,
		Schema: map[string]*schema.Schema{
			"field_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fields": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseCustomField(),
				},
			},
			"include_inactive": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dataSourceGreenhouseCustomFieldsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	fieldType := d.Get("field_type").(string)
	includeInactive := d.Get("include_inactive").(bool)
	list, err := greenhouse.GetAllCustomFields(meta.(*greenhouse.Client), ctx, fieldType, includeInactive)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId("all")
	d.Set("fields", flattenCustomFields(ctx, list))
	return nil
}

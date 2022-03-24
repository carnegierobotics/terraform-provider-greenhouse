package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseEducationDegree() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseEducationDegreeRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseEducationDegreeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllDegrees(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	name := d.Get("name").(string)
	for _, degree := range *list {
		if degree.Name == name {
			d.SetId(strconv.Itoa(degree.Id))
			d.Set("priority", degree.Priority)
			return nil
		}
	}
	return nil
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseEducationDiscipline() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseEducationDisciplineRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseEducationDisciplineRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllDisciplines(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	name := d.Get("name").(string)
	for _, discipline := range *list {
		if *discipline.Name == name {
			d.SetId(strconv.Itoa(*discipline.Id))
			d.Set("priority", discipline.Priority)
			return nil
		}
	}
	return nil
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseDemographicQuestionSet() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDemographicQuestionSetRead,
		Schema:      schemaGreenhouseDemographicQuestionSet(),
	}
}

func dataSourceGreenhouseDemographicQuestionSetRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	title := d.Get("title").(string)
	list, err := greenhouse.GetAllDemographicQuestionSets(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for _, set := range *list {
		if *set.Title == title {
			d.SetId(strconv.Itoa(*set.Id))
			d.Set("active", set.Active)
			d.Set("description", set.Description)
			return nil
		}
	}
	return nil
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseDemographicAnswerOption() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDemographicAnswerOptionRead,
		Schema:      schemaGreenhouseDemographicAnswerOption(),
	}
}

func dataSourceGreenhouseDemographicAnswerOptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("id").(int)
	option, err := greenhouse.GetDemographicAnswerOption(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(option.Id))
	d.Set("active", option.Active)
	d.Set("demographic_question_id", option.DemographicQuestionId)
	d.Set("free_form", option.FreeForm)
	d.Set("name", option.Name)
	d.Set("translations", flattenTranslations(ctx, &option.Translations))
	return nil
}

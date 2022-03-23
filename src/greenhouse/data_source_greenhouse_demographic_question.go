package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseDemographicQuestion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDemographicQuestionRead,
		Schema:      schemaGreenhouseDemographicQuestion(),
	}
}

func dataSourceGreenhouseDemographicQuestionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := d.Get("name").(string)
	list, err := greenhouse.GetAllDemographicQuestions(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for _, question := range *list {
		if question.Name == name {
			d.SetId(strconv.Itoa(question.Id))
      d.Set("active", question.Active)
      d.Set("answer_type", question.AnswerType)
      d.Set("demographic_question_set_id", question.DemographicQuestionSetId)
      d.Set("required", question.Required)
      d.Set("translations", flattenTranslations(ctx, &question.Translations))
			return nil
		}
	}
	return nil
}

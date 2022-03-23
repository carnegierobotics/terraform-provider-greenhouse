package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseDemographicAnswer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseDemographicAnswerRead,
		Schema:      schemaGreenhouseDemographicAnswer(),
	}
}

func dataSourceGreenhouseDemographicAnswerRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("id").(int)
	answer, err := greenhouse.GetDemographicAnswer(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(answer.Id))
	d.Set("application_id", answer.ApplicationId)
	d.Set("created_at", answer.CreatedAt)
	d.Set("demographic_answer_option_id", answer.DemographicAnswerOptionId)
	d.Set("demographic_question_id", answer.DemographicQuestionId)
	d.Set("free_form_text", answer.FreeFormText)
	d.Set("updated_at", answer.UpdatedAt)
	return nil
}

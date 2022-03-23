package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseScorecard() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseScorecardRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseScorecardRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := d.Get("id").(int)
	card, err := greenhouse.GetScorecard(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(card.Id))
	d.Set("application_id", card.ApplicationId)
	d.Set("attributes", flattenScorecardAttributes(ctx, &card.Attributes))
	d.Set("candidate_id", card.CandidateId)
	d.Set("created_at", card.CreatedAt)
	d.Set("interview", card.Interview)
	convertedStep := greenhouse.TypeIdName(*card.InterviewStep)
	d.Set("interview_step", flattenTypeIdName(ctx, &convertedStep))
	d.Set("interviewer", flattenUser(ctx, card.Interviewer))
	d.Set("overall_recommendation", card.OverallRecommendation)
	d.Set("questions", flattenScorecardQuestions(ctx, &card.Questions))
	d.Set("ratings", card.Ratings)
	d.Set("submitted_at", card.SubmittedAt)
	d.Set("submitted_by", flattenUser(ctx, card.SubmittedBy))
	d.Set("updated_at", card.UpdatedAt)
	return nil
}

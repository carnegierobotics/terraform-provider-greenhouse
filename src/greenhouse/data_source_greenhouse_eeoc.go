package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseEEOC() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseEEOCRead,
		Schema:      schemaGreenhouseEEOC(),
	}
}

func dataSourceGreenhouseEEOCRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllEEOC(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	applicationId := d.Get("candidate_id").(int)
	candidateId := d.Get("candidate_id").(int)
	for _, eeoc := range *list {
		if eeoc.CandidateId == candidateId && eeoc.ApplicationId == applicationId {
			d.SetId(fmt.Sprintf("%s-%s", strconv.Itoa(applicationId), strconv.Itoa(candidateId)))
			d.Set("disability_status", flattenEEOCAnswer(ctx, eeoc.DisabilityStatus))
			d.Set("gender", flattenEEOCAnswer(ctx, eeoc.Gender))
			d.Set("race", flattenEEOCAnswer(ctx, eeoc.Race))
			d.Set("submitted_at", eeoc.SubmittedAt)
			d.Set("veteran_status", flattenEEOCAnswer(ctx, eeoc.VeteranStatus))
			return nil
		}
	}
	return nil
}

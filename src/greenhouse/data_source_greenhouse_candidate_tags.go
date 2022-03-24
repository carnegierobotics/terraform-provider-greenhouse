package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseCandidateTags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCandidateTagsRead,
		Schema:      schemaGreenhouseCandidateTags(),
	}
}

func dataSourceGreenhouseCandidateTagsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	candidateId, ok := d.GetOk("candidate_id")
	var id string
	var tags *[]greenhouse.CandidateTag
	var err error
	if ok {
		id = strconv.Itoa(candidateId.(int))
		tags, err = greenhouse.GetTagsForCandidate(meta.(*greenhouse.Client), ctx, candidateId.(int))
	} else {
		id = "all"
		tags, err = greenhouse.GetAllCandidateTags(meta.(*greenhouse.Client), ctx)
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(id)
	d.Set("tags", flattenCandidateTags(ctx, tags))
	return nil
}

package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGreenhouseCandidates() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseCandidatesRead,
		Schema: map[string]*schema.Schema{
			"candidate_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"candidates": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseCandidate(),
				},
			},
			"created_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"job_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"updated_after": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_before": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceGreenhouseCandidatesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	list, err := greenhouse.GetAllCandidates(meta.(*greenhouse.Client), ctx)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	nameList := make([]string, len(*list), len(*list))
	for i, item := range *list {
		nameList[i] = fmt.Sprintf("%s %s", *item.FirstName, *item.LastName)
	}
	d.SetId("all")
	d.Set("candidates", flattenCandidates(ctx, list))
	d.Set("names", nameList)
	return nil
}

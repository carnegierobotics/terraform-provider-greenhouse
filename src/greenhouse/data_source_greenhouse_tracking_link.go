package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseTrackingLink() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseTrackingLinkRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceGreenhouseTrackingLinkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	token := d.Get("token").(string)
	link, err := greenhouse.GetTrackingLinkData(meta.(*greenhouse.Client), ctx, token)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(link.Id))
	d.Set("created_at", link.CreatedAt)
	d.Set("credited_to", flattenUser(ctx, link.CreditedTo))
	d.Set("job_board", flattenJobBoard(ctx, link.JobBoard))
	d.Set("job_id", link.JobId)
	d.Set("job_post_id", link.JobPostId)
	d.Set("related_post_id", link.RelatedPostId)
	d.Set("related_post_type", link.RelatedPostType)
	d.Set("source", flattenSource(ctx, link.Source))
	d.Set("updated_at", link.UpdatedAt)
	return nil
}

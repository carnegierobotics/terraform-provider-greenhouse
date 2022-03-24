package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseOffers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseOffersRead,
		Schema:      schemaGreenhouseOffers(),
	}
}

func dataSourceGreenhouseOffersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	applicationId, ok := d.GetOk("application_id")
	var offers *[]greenhouse.Offer
	var err error
	var id string
	if ok {
		id = strconv.Itoa(applicationId.(int))
		offers, err = greenhouse.GetAllOffersForApplication(meta.(*greenhouse.Client), ctx, applicationId.(int))
	} else {
		id = "all"
		offers, err = greenhouse.GetAllOffers(meta.(*greenhouse.Client), ctx)
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(id)
	d.Set("offers", flattenOffers(ctx, offers))
	return nil
}

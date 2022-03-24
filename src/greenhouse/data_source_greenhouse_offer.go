package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseOffer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseOfferRead,
		Schema:      schemaGreenhouseOffer(),
	}
}

func dataSourceGreenhouseOfferRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	offerId, oidOk := d.GetOk("id")
	applicationId, aidOk := d.GetOk("application_id")
	var offer *greenhouse.Offer
	var err error
	var id int
	if aidOk {
		id = applicationId.(int)
		offer, err = greenhouse.GetCurrentOfferForApplication(meta.(*greenhouse.Client), ctx, applicationId.(int))
	} else if oidOk {
		id = offerId.(int)
		offer, err = greenhouse.GetOffer(meta.(*greenhouse.Client), ctx, id)
	} else {
		return nil
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	for k, v := range flattenOffer(ctx, offer) {
		d.Set(k, v)
	}
	return nil
}

package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseOffer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseOfferCreate,
		ReadContext:   resourceGreenhouseOfferRead,
		UpdateContext: resourceGreenhouseOfferUpdate,
		DeleteContext: resourceGreenhouseOfferDelete,
		Exists:        resourceGreenhouseOfferExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseOffer(),
	}
}

func resourceGreenhouseOfferExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/offers/%d", id))
}

func resourceGreenhouseOfferCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var obj greenhouse.Offer
	id := d.Get("application_id").(int)
	if v, ok := d.Get("created_at").(string); ok && len(v) > 0 {
		obj.CreatedAt = &v
	}
	if v, ok := d.Get("custom_fields").([]interface{}); ok && len(v) > 0 {
		obj.CustomFields = v[0].(map[string]interface{})
	}
	if v, ok := d.Get("start_date").(string); ok && len(v) > 0 {
		obj.StartsAt = &v
	}
	if v, ok := d.Get("sent_at").(string); ok && len(v) > 0 {
		obj.SentAt = &v
	}
	err := greenhouse.UpdateCurrentOffer(meta.(*greenhouse.Client), ctx, id, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	return resourceGreenhouseOfferRead(ctx, d, meta)
}

func resourceGreenhouseOfferRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetOffer(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenOffer(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseOfferUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceGreenhouseOfferCreate(ctx, d, meta)
}

func resourceGreenhouseOfferDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for offers."}}
}

package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseCandidateTag() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseCandidateTagCreate,
		ReadContext:   resourceGreenhouseCandidateTagRead,
		UpdateContext: resourceGreenhouseCandidateTagUpdate,
		DeleteContext: resourceGreenhouseCandidateTagDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseTypeIdName(),
	}
}

func resourceGreenhouseCandidateTagCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.CandidateTag{
		Name: d.Get("name").(string),
	}
	id, err := greenhouse.CreateCandidateTag(meta.(*greenhouse.Client), ctx, &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseUserUpdate(ctx, d, meta)
}

func resourceGreenhouseCandidateTagRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetCandidateTag(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if obj != nil {
		tflog.Debug(ctx, fmt.Sprintf("Could not find tag with id %d", id))
		return nil
	}
	for k, v := range flattenCandidateTag(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseCandidateTagUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for candidate_tag."}}
}

func resourceGreenhouseCandidateTagDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	err = greenhouse.DeleteCandidateTag(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return nil
}

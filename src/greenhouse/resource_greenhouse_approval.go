package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseApproval() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseApprovalCreate,
		ReadContext:   resourceGreenhouseApprovalRead,
		UpdateContext: resourceGreenhouseApprovalUpdate,
		DeleteContext: resourceGreenhouseApprovalDelete,
		Exists:        resourceGreenhouseApprovalExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseApprovalImport,
		},
		Schema: schemaGreenhouseApproval(),
	}
}

func resourceGreenhouseApprovalExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/approval_flows/%d", id))
}

func resourceGreenhouseApprovalCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	jobId := d.Get("job_id").(int)
	obj, diagErr := createApprovalObj(ctx, d)
	if diagErr != nil {
		return diagErr
	}
	id, err := greenhouse.CreateReplaceApprovalFlow(meta.(*greenhouse.Client), ctx, jobId, obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	if d.HasChange("request_approval") {
		if v, ok := d.Get("request_approval").(bool); ok && v {
			err := greenhouse.RequestApprovals(meta.(*greenhouse.Client), ctx, id)
			if err != nil {
				return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
			}
		}
	}
	return resourceGreenhouseApprovalRead(ctx, d, meta)
}

func resourceGreenhouseApprovalRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.RetrieveApprovalFlow(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenApproval(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseApprovalUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceGreenhouseApprovalRead(ctx, d, meta)
}

func resourceGreenhouseApprovalDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for approvals."}}
}

func resourceGreenhouseApprovalImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseApprovalRead)
}

func createApprovalObj(ctx context.Context, d *schema.ResourceData) (*greenhouse.Approval, diag.Diagnostics) {
	var obj greenhouse.Approval
	if v, ok := d.Get("approval_type").(string); ok && len(v) > 0 {
		obj.ApprovalType = &v
	}
	if v, ok := d.Get("approver_groups").([]interface{}); ok && len(v) > 0 {
		list, err := inflateApproverGroups(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.ApproverGroups = *list
	}
	if v, ok := d.Get("offer_id").(int); ok {
		obj.OfferId = &v
	}
	if v, ok := d.Get("sequential").(bool); ok {
		obj.Sequential = &v
	}
	return &obj, nil
}

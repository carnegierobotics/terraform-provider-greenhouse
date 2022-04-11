package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseJobPost() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseJobPostCreate,
		ReadContext:   resourceGreenhouseJobPostRead,
		UpdateContext: resourceGreenhouseJobPostUpdate,
		DeleteContext: resourceGreenhouseJobPostDelete,
		Exists:        resourceGreenhouseJobPostExists,
		Importer: &schema.ResourceImporter{
			StateContext: resourceGreenhouseJobPostImport,
		},
		Schema: schemaGreenhouseJobPost(),
	}
}

func resourceGreenhouseJobPostExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/job_posts/%d", id))
}

func resourceGreenhouseJobPostCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Create is not supported for job_posts."}}
}

func resourceGreenhouseJobPostRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetJobPost(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenJobPost(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseJobPostUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	if d.HasChanges("live") {
		status := "offline"
		if v, ok := d.Get("live").(bool); ok && v {
			status = "live"
		}
		err := greenhouse.UpdateJobPostStatus(meta.(*greenhouse.Client), ctx, id, status)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	var obj greenhouse.JobPost
	if v, ok := d.Get("location").(string); ok && len(v) > 0 {
		obj.Location = &v
	}
	if v, ok := d.Get("title").(string); ok && len(v) > 0 {
		obj.Title = &v
	}
	if v, ok := d.Get("content").(string); ok && len(v) > 0 {
		obj.Content = &v
	}
	err = greenhouse.UpdateJobPost(meta.(*greenhouse.Client), ctx, id, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return resourceGreenhouseJobPostRead(ctx, d, meta)
}

func resourceGreenhouseJobPostDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for job_posts."}}
}

func resourceGreenhouseJobPostImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return importByRead(ctx, d, meta, resourceGreenhouseJobPostRead)
}

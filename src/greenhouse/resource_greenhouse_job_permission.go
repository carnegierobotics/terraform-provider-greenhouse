package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseJobPermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseJobPermissionCreate,
		ReadContext:   resourceGreenhouseJobPermissionRead,
		UpdateContext: resourceGreenhouseJobPermissionUpdate,
		DeleteContext: resourceGreenhouseJobPermissionDelete,
		Exists:        resourceGreenhouseJobPermissionExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseUserPermission(),
	}
}

func resourceGreenhouseJobPermissionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/xxx/%d", id))
}

func resourceGreenhouseJobPermissionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  var obj greenhouse.UserPermission
  if v, ok := d.Get("job_id").(int); ok {
    obj.JobId = v
  }
  if v, ok := d.Get("user_role_id").(int); ok {
    obj.UserRoleId = v
  }
  if v, ok := d.Get("user_id").(int); ok {
    id, err := greenhouse.CreateJobPermission(meta.(*greenhouse.Client), ctx, v, &obj)
    if err != nil {
      return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
    }
    d.SetId(strconv.Itoa(id))
    return resourceGreenhouseJobPermissionRead(ctx, d, meta)
  }
  return diag.Diagnostics{{Severity: diag.Error, Summary: "Could not create permission."}}
}

func resourceGreenhouseJobPermissionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  userId, ok := d.Get("user_id").(int)
  if !ok {
    return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting user_id."}}
  }
	obj, err := greenhouse.GetJobPermission(meta.(*greenhouse.Client), ctx, userId, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.Set("job_id", obj.JobId)
  d.Set("user_role_id", obj.UserRoleId)
	return nil
}

func resourceGreenhouseJobPermissionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for xxx."}}
}

func resourceGreenhouseJobPermissionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  permId, err := strconv.Atoi(d.Id())
  if err != nil {
    return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
  }
  jobId, ok := d.Get("job_id").(int)
  if !ok {
    return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting job_id."}} 
  }
  err = greenhouse.DeleteJobPermission(meta.(*greenhouse.Client), ctx, jobId, permId)
  if err != nil {
    return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
  }
  d.SetId("")
  return nil
}

package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseFutureJobPermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseFutureJobPermissionCreate,
		ReadContext:   resourceGreenhouseFutureJobPermissionRead,
		UpdateContext: resourceGreenhouseFutureJobPermissionUpdate,
		DeleteContext: resourceGreenhouseFutureJobPermissionDelete,
		Exists:        resourceGreenhouseFutureJobPermissionExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseFutureJobPermission(),
	}
}

func resourceGreenhouseFutureJobPermissionExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/xxx/%d", id))
}

func resourceGreenhouseFutureJobPermissionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  var obj greenhouse.FutureJobPermission
  if v, ok := d.Get("department_id").(int); ok {
    obj.DepartmentId = v
  }
  if v, ok := d.Get("external_department_id").(string); ok && len(v) > 0 {
    obj.ExternalDepartmentId = v
  }
  if v, ok := d.Get("external_office_id").(string); ok {
    obj.ExternalOfficeId = v
  }
  if v, ok := d.Get("office_id").(int); ok {
    obj.OfficeId = v
  }
  if v, ok := d.Get("user_role_id").(int); ok {
    obj.UserRoleId = v
  }
  if v, ok := d.Get("user_id").(int); ok {
    id, err := greenhouse.CreateFutureJobPermission(meta.(*greenhouse.Client), ctx, v, &obj)
    if err != nil {
      return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
    }
    d.SetId(strconv.Itoa(id))
    return resourceGreenhouseFutureJobPermissionRead(ctx, d, meta)
  }
  return diag.Diagnostics{{Severity: diag.Error, Summary: "Could not create permission."}}
}

func resourceGreenhouseFutureJobPermissionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  userId, ok := d.Get("user_id").(int)
  if !ok {
    return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting user_id."}}
  }
	obj, err := greenhouse.GetFutureJobPermission(meta.(*greenhouse.Client), ctx, userId, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
  d.Set("department_id", obj.DepartmentId)
  d.Set("external_department_id", obj.ExternalDepartmentId)
  d.Set("external_office_id", obj.ExternalOfficeId)
	d.Set("office_id", obj.OfficeId)
  d.Set("user_role_id", obj.UserRoleId)
	return nil
}

func resourceGreenhouseFutureJobPermissionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Update is not supported for xxx."}}
}

func resourceGreenhouseFutureJobPermissionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
  permId, err := strconv.Atoi(d.Id())
  if err != nil {
    return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
  }
  jobId, ok := d.Get("job_id").(int)
  if !ok {
    return diag.Diagnostics{{Severity: diag.Error, Summary: "Error getting job_id."}}
  }
  err = greenhouse.DeleteFutureJobPermission(meta.(*greenhouse.Client), ctx, jobId, permId)
  if err != nil {
    return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
  }
  d.SetId("")
  return nil
}

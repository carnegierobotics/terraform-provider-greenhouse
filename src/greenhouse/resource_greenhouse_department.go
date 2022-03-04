package greenhouse

import (
  "context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

type ReadFunc func(d *schema.ResourceData, m interface{}) error

func resourceGreenhouseDepartment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseDepartmentCreate,
		ReadContext:   resourceGreenhouseDepartmentRead,
		UpdateContext: resourceGreenhouseDepartmentUpdate,
		DeleteContext: resourceGreenhouseDepartmentDelete,
		Exists: resourceGreenhouseDepartmentExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseDepartment(),
	}
}

func resourceGreenhouseDepartmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "departments", id, context.TODO())
}

func resourceGreenhouseDepartmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.DepartmentCreateInfo{
		Name:     d.Get("name").(string),
		ParentId: d.Get("parent_id").(int),
	}
	id, err := greenhouse.CreateDepartment(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseDepartmentRead(ctx, d, meta)
}

func resourceGreenhouseDepartmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	obj, err := greenhouse.GetDepartment(meta.(*greenhouse.Client), id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	d.Set("name", obj.Name)
	d.Set("parent_id", obj.ParentId)
	d.Set("child_ids", obj.ChildIds)
	return nil
}

func resourceGreenhouseDepartmentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	updateObject := greenhouse.DepartmentUpdateInfo{
		Name: d.Get("name").(string),
	}
	err = greenhouse.UpdateDepartment(meta.(*greenhouse.Client), id, &updateObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	return resourceGreenhouseDepartmentRead(ctx, d, meta)
}

func resourceGreenhouseDepartmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for departments."},}
}

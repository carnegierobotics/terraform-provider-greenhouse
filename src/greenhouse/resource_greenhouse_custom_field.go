package greenhouse

import (
	"context"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseCustomField() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseCustomFieldCreate,
		ReadContext:   resourceGreenhouseCustomFieldRead,
		UpdateContext: resourceGreenhouseCustomFieldUpdate,
		DeleteContext: resourceGreenhouseCustomFieldDelete,
		Exists:        resourceGreenhouseCustomFieldExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseCustomField(),
	}
}

func resourceGreenhouseCustomFieldExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/custom_field/%d", id))
}

func resourceGreenhouseCustomFieldCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	obj, diagErr := createCustomFieldObject(ctx, d)
	if diagErr != nil {
		return diagErr
	}
	id, err := greenhouse.CreateCustomField(meta.(*greenhouse.Client), ctx, obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(id))
	return resourceGreenhouseCustomFieldUpdate(ctx, d, meta)
}

func resourceGreenhouseCustomFieldRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, err := greenhouse.GetCustomField(meta.(*greenhouse.Client), ctx, id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	for k, v := range flattenCustomField(ctx, obj) {
		d.Set(k, v)
	}
	return nil
}

func resourceGreenhouseCustomFieldUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	obj, diagErr := createCustomFieldObject(ctx, d)
	if diagErr != nil {
		return diagErr
	}
	if d.HasChange("template_token_string") {
		if v, ok := d.Get("template_token_string").(string); ok && len(v) > 0 {
			obj.TemplateTokenString = v
		}
	}
	err = greenhouse.UpdateCustomField(meta.(*greenhouse.Client), ctx, id, obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	return resourceGreenhouseCustomFieldRead(ctx, d, meta)
}

func resourceGreenhouseCustomFieldDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for custom_field."}}
}

func createCustomFieldObject(ctx context.Context, d *schema.ResourceData) (*greenhouse.CustomField, diag.Diagnostics) {
	var obj greenhouse.CustomField
	if v, ok := d.Get("api_only").(bool); ok {
		obj.ApiOnly = v
	}
	if v, ok := d.Get("custom_field_options").([]interface{}); ok && len(v) > 0 {
		list, err := inflateCustomFieldOptions(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.CustomFieldOptions = *list
	}
	if v, ok := d.Get("department_ids").([]interface{}); ok && len(v) > 0 {
		obj.DepartmentIds = *sliceItoSliceD(&v)
	}
	if v, ok := d.Get("description").(string); ok && len(v) > 0 {
		obj.Description = v
	}
	if v, ok := d.Get("expose_in_job_board_api").(bool); ok {
		obj.ExposeInJobBoardAPI = v
	}
	if v, ok := d.Get("field_type").(string); ok && len(v) > 0 {
		obj.FieldType = v
	}
	if v, ok := d.Get("generate_email_token").(bool); ok {
		obj.GenerateEmailToken = v
	}
	if v, ok := d.Get("name").(string); ok && len(v) > 0 {
		obj.Name = v
	}
	if v, ok := d.Get("office_ids").([]interface{}); ok && len(v) > 0 {
		obj.OfficeIds = *sliceItoSliceD(&v)
	}
	if v, ok := d.Get("private").(bool); ok {
		obj.Private = v
	}
	if v, ok := d.Get("require_approval").(bool); ok {
		obj.RequireApproval = v
	}
	if v, ok := d.Get("required").(bool); ok {
		obj.Required = v
	}
	if v, ok := d.Get("trigger_new_version").(bool); ok {
		obj.TriggerNewVersion = v
	}
	if v, ok := d.Get("value_type").(string); ok && len(v) > 0 {
		obj.ValueType = v
	}
	return &obj, nil
}

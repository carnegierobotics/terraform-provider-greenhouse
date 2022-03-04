package greenhouse

import (
  "context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseOffice() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseOfficeCreate,
		ReadContext:   resourceGreenhouseOfficeRead,
		UpdateContext: resourceGreenhouseOfficeUpdate,
		DeleteContext: resourceGreenhouseOfficeDelete,
		Exists: resourceGreenhouseOfficeExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseOffice(),
	}
}

func resourceGreenhouseOfficeExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), "offices", id, context.TODO())
}

func resourceGreenhouseOfficeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	createObject := greenhouse.OfficeCreateInfo{
		Name:                 d.Get("name").(string),
		Location:             d.Get("location.name").(string),
		PrimaryContactUserId: d.Get("primary_contact_user_id").(int),
		ParentId:             d.Get("parent_id").(int),
	}
	id, err := greenhouse.CreateOffice(meta.(*greenhouse.Client), &createObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	strId := strconv.Itoa(id)
	d.SetId(strId)
	return resourceGreenhouseOfficeRead(ctx, d, meta)
}

func resourceGreenhouseOfficeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	obj, err := greenhouse.GetOffice(meta.(*greenhouse.Client), id)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	d.Set("name", obj.Name)
	d.Set("location", flattenLocation(&obj.Location))
	d.Set("primary_contact_user_id", obj.PrimaryContactUserId)
	d.Set("parent_id", obj.ParentId)
	d.Set("child_ids", obj.ChildIds)
	return nil
}

func resourceGreenhouseOfficeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	updateObject := greenhouse.OfficeUpdateInfo{
		Name:     d.Get("name").(string),
		Location: d.Get("location.name").(string),
	}
	err = greenhouse.UpdateOffice(meta.(*greenhouse.Client), id, &updateObject)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()},}
	}
	return resourceGreenhouseOfficeRead(ctx, d, meta)
}

func resourceGreenhouseOfficeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for offices."},}
}

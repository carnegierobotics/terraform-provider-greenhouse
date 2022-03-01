package greenhouse

import (
  "github.com/carnegierobotics/greenhouse-client-go/greenhouse"
  "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type ReadFunc func(d *schema.ResourceData, m interface{}) error

func resourceGreenhouseDepartment() *schema.Resource {
  return &schema.Resource{
    Create: resourceGreenhouseDepartmentCreate,
    Read: resourceGreenhouseDepartmentRead,
    Update: resourceGreenhouseDepartmentUpdate,
    Exists: resourceGreenhouseDepartmentExists,
    Importer: &schema.ResourceImporter{
      State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
        return []*schema.ResourceData{d}, nil
      },
    },
    Schema: map[string]*schema.Schema {
      "id": {
        Type: schema.TypeInt,
        Required: false,
        Computed: true,
      },
      "name": {
        Type: schema.TypeString,
        Required: true,
      },
      "parent_id": {
        Type: schema.TypeInt,
        Optional: true,
      },
      "child_ids": {
        Type: schema.TypeSet,
        Optional: true,
        Elem: &schema.Schema {
          Type: schema.TypeInt,
        },
      },
      /* Not in our product tier
      "parent_department_external_id": {
        Type: schema.TypeString,
        Optional: true,
        Computed: true,
      },
      "child_department_external_ids": {
        Type: schema.TypeSet,
        Optional: true,
        Computed: true,
        Elem: &schema.Schema {
          Type: schema.TypeString,
        }
      },
      "external_id": {
        Type: schema.TypeString,
        Optional: true,
      }
      */
    },
  }
}

func resourceGreenhouseDepartmentObject(d *schema.ResourceData) *greenhouse.Department {
  return &greenhouse.Department {
    Name: d.Get("name").(string),
    ParentId: d.Get("parent_id").(int),
    ChildIds: ConvertSliceInterfaceInt(d.Get("child_ids").(*schema.Set).List()),
  }
}

func resourceGreenhouseDepartmentExists(d *schema.ResourceData, meta interface{}) (bool, error) {
  id := d.Get("id").(int)
  return greenhouse.Exists(meta.(*greenhouse.Client), "departments", id)
}

func resourceGreenhouseDepartmentCreate(d *schema.ResourceData, meta interface{}) error {
  createObject := greenhouse.DepartmentCreateInfo{
    Name: d.Get("name").(string),
    ParentId: d.Get("parent_id").(int),
  }
  err := greenhouse.CreateDepartment(meta.(*greenhouse.Client), &createObject)
  if err != nil {
    return err
  }
  d.SetId(createObject.Name)
  return resourceGreenhouseDepartmentRead(d, meta)
}

func resourceGreenhouseDepartmentRead(d *schema.ResourceData, meta interface{}) error {
  id := d.Get("id").(int)
  obj, err := greenhouse.GetDepartment(meta.(*greenhouse.Client), id)
  if err != nil {
    return err
  }
  d.Set("name", obj.Name)
  d.Set("parent_id", obj.ParentId)
  d.Set("child_ids", obj.ChildIds)
  return nil
}

func resourceGreenhouseDepartmentUpdate(d *schema.ResourceData, meta interface{}) error {
  id := d.Get("id").(int)
  updateObject := greenhouse.DepartmentUpdateInfo{
    Name: d.Get("name").(string),
  }
  err := greenhouse.UpdateDepartment(meta.(*greenhouse.Client), id, &updateObject)
  if err != nil {
    return err
  }
  return resourceGreenhouseDepartmentRead(d, meta)
}

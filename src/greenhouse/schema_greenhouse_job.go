package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJob() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"anywhere": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"closed_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"confidential": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"copied_from_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"department_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"departments": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"hiring_team": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseHiringTeam(),
			},
		},
		"how_to_sell_this_job": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"is_template": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"job_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"job_post_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		/* TODO
		"keyed_custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeMap,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		*/
		"notes": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"number_of_openings": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"office_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"offices": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffice(),
			},
		},
		"opened_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"opening_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"openings": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseJobOpening(),
			},
		},
		"requisition_id": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"team_and_responsibilities": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"template_job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateJobs(ctx context.Context, source *[]interface{}) (*[]greenhouse.Job, diag.Diagnostics) {
	list := make([]greenhouse.Job, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateJob(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateJob(ctx context.Context, source *map[string]interface{}) (*greenhouse.Job, diag.Diagnostics) {
	var obj greenhouse.Job
	if v, ok := (*source)["closed_at"].(string); ok && len(v) > 0 {
		obj.ClosedAt = v
	}
	if v, ok := (*source)["confidential"].(bool); ok {
		obj.Confidential = v
	}
	if v, ok := (*source)["copied_from_id"].(int); ok {
		obj.CopiedFromId = v
	}
	if v, ok := (*source)["created_at"].(string); ok && len(v) > 0 {
		obj.CreatedAt = v
	}
	/*
	  if v, ok := (*source)["custom_fields"].(map[string]string); ok && len(v) > 0 {
	    n := make(map[string]interface{})
	    for k, v := range v {
	      n[k] = v.(interface{})
	    }
	    obj.CustomFields = n
	  }
	*/
	if v, ok := (*source)["departments"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateDepartments(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Departments = *list
	}
	if v, ok := (*source)["hiring_team"].([]interface{}); ok && len(v) > 0 {
		team := v[0].(map[string]interface{})
		teamMap, err := inflateHiringTeams(ctx, &team)
		if err != nil {
			return nil, err
		}
		obj.HiringTeam = *teamMap
	}
	if v, ok := (*source)["is_template"].(bool); ok {
		obj.IsTemplate = v
	}
	if v, ok := (*source)["name"].(string); ok && len(v) > 0 {
		obj.Name = v
	}
	if v, ok := (*source)["notes"].(string); ok && len(v) > 0 {
		obj.Notes = v
	}
	if v, ok := (*source)["offices"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateOffices(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Offices = *list
	}
	if v, ok := (*source)["opened_at"].(string); ok && len(v) > 0 {
		obj.OpenedAt = v
	}
	if v, ok := (*source)["openings"].([]interface{}); ok && len(v) > 0 {
		list, err := inflateJobOpenings(ctx, &v)
		if err != nil {
			return nil, err
		}
		obj.Openings = *list
	}
	if v, ok := (*source)["requisition_id"].(string); ok && len(v) > 0 {
		obj.RequisitionId = v
	}
	if v, ok := (*source)["status"].(string); ok && len(v) > 0 {
		obj.Status = v
	}
	if v, ok := (*source)["updated_at"].(string); ok && len(v) > 0 {
		obj.UpdatedAt = v
	}
	return &obj, nil
}

func flattenJobs(ctx context.Context, list *[]greenhouse.Job) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenJob(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenJob(ctx context.Context, item *greenhouse.Job) map[string]interface{} {
	job := make(map[string]interface{})
	job["closed_at"] = item.ClosedAt
	job["confidential"] = item.Confidential
	job["copied_from_id"] = item.CopiedFromId
	job["created_at"] = item.CreatedAt
	job["custom_fields"] = item.CustomFields
	job["departments"] = flattenDepartments(ctx, &item.Departments)
	job["hiring_team"] = flattenHiringTeam(ctx, &item.HiringTeam)
	job["is_template"] = item.IsTemplate
	job["job_name"] = item.Name
	//job["keyed_custom_fields"] = flattenKeyedCustomFields(ctx, &item.KeyedCustomFields)
	job["notes"] = item.Notes
	job["offices"] = flattenOffices(ctx, &item.Offices)
	job["opened_at"] = item.OpenedAt
	job["openings"] = flattenJobOpenings(ctx, &item.Openings)
	job["requisition_id"] = item.RequisitionId
	job["status"] = item.Status
	job["updated_at"] = item.UpdatedAt
	return job
}

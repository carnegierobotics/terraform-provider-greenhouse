package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJob() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"requisition_id": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"notes": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"anywhere": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"team_and_responsibilities": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"how_to_sell_this_job": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"confidential": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"opened_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"closed_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"is_template": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"copied_from_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"department_id": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"departments": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"office_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"offices": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseOffice(),
			},
		},
		"custom_fields": {
			Type:     schema.TypeMap,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		/*
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
		"hiring_team": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseHiringTeam(),
			},
		},
		"number_of_openings": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"template_job_id": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"openings": {
			Type:     schema.TypeSet,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseJobOpening(),
			},
		},
		"opening_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
	}
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

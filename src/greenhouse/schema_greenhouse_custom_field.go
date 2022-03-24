package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func schemaGreenhouseCustomField() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"api_only": {
			Type:        schema.TypeBool,
			Description: "Updates to this field may only be made via Harvest.",
			Optional:    true,
			Default:     false,
		},
		"custom_field_options": {
			Type:        schema.TypeList,
			Description: "For single_select and multi_select field_types, this is the list of options for that select.",
			Optional:    true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseCustomFieldOption(),
			},
		},
		"department_ids": {
			Type:        schema.TypeSet,
			Description: "The custom field is only displayed for objects in these departments.",
			Optional:    true,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
		},
		"departments": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDepartment(),
			},
		},
		"description": {
			Type:        schema.TypeString,
			Description: "The field's description in Greenhouse",
			Optional:    true,
		},
		"expose_in_job_board_api": {
			Type:        schema.TypeBool,
			Description: "This field and its value are provided in the Job Board API response.",
			Optional:    true,
			Default:     false,
		},
		"field_type": {
			Type:         schema.TypeString,
			Description:  "The type of the field.",
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"job", "candidate", "application", "offer", "opening", "rejection_question", "referral_question"}, false),
		},
		"generate_email_token": {
			Type:        schema.TypeBool,
			Description: "Generate a default template_token_string for the new Custom Field.",
			Optional:    true,
		},
		"name": {
			Type:        schema.TypeString,
			Description: "The field's name in Greenhouse",
			Required:    true,
		},
		"name_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"office_ids": {
			Type:        schema.TypeSet,
			Description: "The custom field is only displayed for objects in these offices.",
			Optional:    true,
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
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"private": {
			Type:        schema.TypeBool,
			Description: "Denotes a private field in Greenhouse.",
			Optional:    true,
			Default:     false,
		},
		"require_approval": {
			Type:        schema.TypeBool,
			Description: "Changes to this field trigger re-approval.",
			Optional:    true,
			Default:     false,
		},
		"required": {
			Type:        schema.TypeBool,
			Description: "Denotes a required field in Greenhouse.",
			Optional:    true,
			Default:     false,
		},
		"template_token_string": {
			Type:        schema.TypeString,
			Description: "The template token used in email and offer document templates.",
			Optional:    true,
			Computed:    true,
		},
		"trigger_new_version": {
			Type:        schema.TypeBool,
			Description: "Changes to this field trigger creation of a new offer version.",
			Optional:    true,
			Default:     false,
		},
		"value_type": {
			Type:         schema.TypeString,
			Description:  "The type of the value.",
			Required:     true,
			ValidateFunc: validation.StringInSlice([]string{"short_text", "long_text", "yes_no", "single_select", "multi_select", "currency", "currency_range", "number", "number_range", "date", "url", "user"}, false),
		},
	}
}

func flattenCustomFields(ctx context.Context, list *[]greenhouse.CustomField) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenCustomField(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenCustomField(ctx context.Context, item *greenhouse.CustomField) map[string]interface{} {
	fields := make(map[string]interface{})
	fields["active"] = item.Active
	fields["api_only"] = item.ApiOnly
	fields["custom_field_options"] = flattenCustomFieldOptions(ctx, &item.CustomFieldOptions)
	fields["departments"] = flattenDepartments(ctx, &item.Departments)
	fields["description"] = item.Description
	fields["expose_in_job_board_api"] = item.ExposeInJobBoardAPI
	fields["field_type"] = item.FieldType
	fields["name"] = item.Name
	fields["name_key"] = item.NameKey
	fields["offices"] = flattenOffices(ctx, &item.Offices)
	fields["priority"] = item.Priority
	fields["private"] = item.Private
	fields["require_approval"] = item.RequireApproval
	fields["required"] = item.Required
	fields["template_token_string"] = item.TemplateTokenString
	fields["trigger_new_version"] = item.TriggerNewVersion
	fields["value_type"] = item.ValueType
	return fields
}

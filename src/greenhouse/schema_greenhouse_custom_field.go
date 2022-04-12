/*
Copyright 2021-2022
Carnegie Robotics, LLC
4501 Hatfield Street, Pittsburgh, PA 15201
https://www.carnegierobotics.com
All rights reserved.

This file is part of terraform-provider-greenhouse.

terraform-provider-greenhouse is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

terraform-provider-greenhouse is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with terraform-provider-greenhouse. If not, see <https://www.gnu.org/licenses/>.
*/
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
			Type:        schema.TypeList,
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
			Type:        schema.TypeList,
			Description: "The custom field is only displayed for objects in these offices.",
			Optional:    true,
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
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"private": {
			Type:        schema.TypeBool,
			Description: "Denotes a private field in Greenhouse.",
			Optional:    true,
		},
		"require_approval": {
			Type:        schema.TypeBool,
			Description: "Changes to this field trigger re-approval.",
			Optional:    true,
		},
		"required": {
			Type:        schema.TypeBool,
			Description: "Denotes a required field in Greenhouse.",
			Optional:    true,
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
	if v := item.Active; v != nil {
		fields["active"] = *v
	}
	if v := item.ApiOnly; v != nil {
		fields["api_only"] = *v
	}
	if v := item.CustomFieldOptions; len(v) > 0 {
		fields["custom_field_options"] = flattenCustomFieldOptions(ctx, &v)
	}
	if v := item.Departments; len(v) > 0 {
		fields["departments"] = flattenDepartments(ctx, &v)
	}
	if v := item.Description; v != nil {
		fields["description"] = *v
	}
	if v := item.ExposeInJobBoardAPI; v != nil {
		fields["expose_in_job_board_api"] = *v
	}
	if v := item.FieldType; v != nil {
		fields["field_type"] = *v
	}
	if v := item.Name; v != nil {
		fields["name"] = *v
	}
	if v := item.NameKey; v != nil {
		fields["name_key"] = *v
	}
	if v := item.Offices; len(v) > 0 {
		fields["offices"] = flattenOffices(ctx, &v)
	}
	if v := item.Priority; v != nil {
		fields["priority"] = *v
	}
	if v := item.Private; v != nil {
		fields["private"] = *v
	}
	if v := item.RequireApproval; v != nil {
		fields["require_approval"] = *v
	}
	if v := item.Required; v != nil {
		fields["required"] = *v
	}
	if v := item.TemplateTokenString; v != nil {
		fields["template_token_string"] = *v
	}
	if v := item.TriggerNewVersion; v != nil {
		fields["trigger_new_version"] = *v
	}
	if v := item.ValueType; v != nil {
		fields["value_type"] = *v
	}
	return fields
}

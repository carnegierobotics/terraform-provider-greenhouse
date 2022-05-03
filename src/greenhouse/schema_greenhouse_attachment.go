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
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseAttachment() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"content": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"content_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"filename": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"visibility": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func inflateAttachments(ctx context.Context, source *[]interface{}) (*[]greenhouse.Attachment, diag.Diagnostics) {
	list := make([]greenhouse.Attachment, len(*source), len(*source))
	for i, item := range *source {
		itemMap := item.(map[string]interface{})
		obj, err := inflateAttachment(ctx, &itemMap)
		if err != nil {
			return nil, err
		}
		list[i] = *obj
	}
	return &list, nil
}

func inflateAttachment(ctx context.Context, source *map[string]interface{}) (*greenhouse.Attachment, diag.Diagnostics) {
	var obj greenhouse.Attachment
	if v, ok := (*source)["content"].(string); ok && len(v) > 0 {
		obj.Content = &v
	}
	if v, ok := (*source)["content_type"].(string); ok && len(v) > 0 {
		obj.ContentType = &v
	}
	if v, ok := (*source)["filename"].(string); ok && len(v) > 0 {
		obj.ContentType = &v
	}
	if v, ok := (*source)["type"].(string); ok && len(v) > 0 {
		obj.Type = &v
	}
	if v, ok := (*source)["url"].(string); ok && len(v) > 0 {
		obj.Url = &v
	}
	if v, ok := (*source)["visibility"].(string); ok && len(v) > 0 {
		obj.Visibility = &v
	}
	return &obj, nil
}

func flattenAttachments(ctx context.Context, list *[]greenhouse.Attachment) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenAttachment(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0)
}

func flattenAttachment(ctx context.Context, item *greenhouse.Attachment) map[string]interface{} {
	attachment := make(map[string]interface{})
	if v := item.Filename; v != nil {
		attachment["filename"] = *v
	}
	if v := item.Type; v != nil {
		attachment["type"] = *v
	}
	if v := item.Url; v != nil {
		attachment["url"] = *v
	}
	return attachment
}

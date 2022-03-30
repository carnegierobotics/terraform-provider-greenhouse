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
    obj.Content = v
  }
  if v, ok := (*source)["content_type"].(string); ok && len(v) > 0 {
    obj.ContentType = v
  }
  if v, ok := (*source)["filename"].(string); ok && len(v) > 0 {
    obj.ContentType = v
  }
  if v, ok := (*source)["type"].(string); ok && len(v) > 0 {
    obj.Type = v
  }
  if v, ok := (*source)["url"].(string); ok && len(v) > 0 {
    obj.Url = v
  }
  if v, ok := (*source)["visibility"].(string); ok && len(v) > 0 {
    obj.Visibility = v
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
	attachment["filename"] = item.Filename
	attachment["type"] = item.Type
	attachment["url"] = item.Url
	return attachment
}

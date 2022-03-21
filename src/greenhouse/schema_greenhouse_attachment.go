package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
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

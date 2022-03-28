package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func schemaGreenhouseJobBoard() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"company_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"url_token": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func flattenJobBoard(ctx context.Context, item *greenhouse.JobBoard) map[string]interface{} {
	board := make(map[string]interface{})
	board["company_name"] = item.CompanyName
	board["id"] = strconv.Itoa(item.Id)
	board["url_token"] = item.UrlToken
	return board
}

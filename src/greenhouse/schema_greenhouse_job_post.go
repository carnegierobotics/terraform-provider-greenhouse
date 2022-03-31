package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func schemaGreenhouseJobPost() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"content": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"demographic_question_set_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"external": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"first_published_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"internal": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"internal_content": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"job_id": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"live": {
			Type:     schema.TypeBool,
			Optional: true,
		},
    "location": {
      Type: schema.TypeString,
      Optional: true,
    },
		"questions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: schemaGreenhouseDemographicQuestion(),
			},
		},
		"updated_at": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func flattenJobPosts(ctx context.Context, list *[]greenhouse.JobPost) []interface{} {
	if list != nil {
		flatList := make([]interface{}, len(*list), len(*list))
		for i, item := range *list {
			flatList[i] = flattenJobPost(ctx, &item)
		}
		return flatList
	}
	return make([]interface{}, 0, 0)
}

func flattenJobPost(ctx context.Context, item *greenhouse.JobPost) map[string]interface{} {
	post := make(map[string]interface{})
	post["active"] = item.Active
	post["content"] = item.Content
	post["created_at"] = item.CreatedAt
	post["demographic_question_set_id"] = item.DemographicQuestionSetId
	post["external"] = item.External
	post["first_published_at"] = item.FirstPublishedAt
	post["internal"] = item.Internal
	post["internal_content"] = item.InternalContent
	post["job_id"] = item.JobId
	post["live"] = item.Live
	post["questions"] = flattenDemographicQuestions(ctx, &item.Questions)
  post["title"] = item.Title
	post["updated_at"] = item.UpdatedAt
	return post
}

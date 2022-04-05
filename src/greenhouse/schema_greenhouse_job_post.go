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
			Type:     schema.TypeString,
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
	if v := item.Active; v != nil {
		post["active"] = *v
	}
	if v := item.Content; v != nil {
		post["content"] = *v
	}
	if v := item.CreatedAt; v != nil {
		post["created_at"] = *v
	}
	if v := item.DemographicQuestionSetId; v != nil {
		post["demographic_question_set_id"] = *v
	}
	if v := item.External; v != nil {
		post["external"] = *v
	}
	if v := item.FirstPublishedAt; v != nil {
		post["first_published_at"] = *v
	}
	if v := item.Internal; v != nil {
		post["internal"] = *v
	}
	if v := item.InternalContent; v != nil {
		post["internal_content"] = *v
	}
	if v := item.JobId; v != nil {
		post["job_id"] = *v
	}
	if v := item.Live; v != nil {
		post["live"] = *v
	}
	if v := item.Questions; len(v) > 0 {
		post["questions"] = flattenDemographicQuestions(ctx, &v)
	}
	if v := item.Title; v != nil {
		post["title"] = *v
	}
	if v := item.UpdatedAt; v != nil {
		post["updated_at"] = *v
	}
	return post
}

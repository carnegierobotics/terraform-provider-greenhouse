package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func dataSourceGreenhouseJobPosts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGreenhouseJobPostsRead,
		Schema: map[string]*schema.Schema{
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"job_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"live": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"posts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: schemaGreenhouseJobPost(),
				},
			},
		},
	}
}

func dataSourceGreenhouseJobPostsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	jobId, ok := d.GetOk("job_id")
	var posts *[]greenhouse.JobPost
	var err error
	var id string
	if ok {
		id = strconv.Itoa(jobId.(int))
		posts, err = greenhouse.GetAllJobPostsForJob(meta.(*greenhouse.Client), ctx, jobId.(int))
	} else {
		id = "all"
		live := d.Get("live").(bool)
		active := d.Get("active").(bool)
		posts, err = greenhouse.GetAllJobPosts(meta.(*greenhouse.Client), ctx, live, active)
	}
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(id)
	d.Set("posts", flattenJobPosts(ctx, posts))
	return nil
}

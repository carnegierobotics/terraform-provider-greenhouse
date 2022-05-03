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

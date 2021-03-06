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
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"harvest_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_HARVEST_TOKEN", nil),
				Description: "The token to use for the Greenhouse Harvest API.",
			},
			"harvest_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_HARVEST_URL", "https://harvest.greenhouse.io"),
				Description: "The URL for Greenhouse's Harvest API.",
			},
			"jobs_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_JOBS_TOKEN", nil),
				Description: "The token to use for the Greenhouse Jobs API.",
			},
			"jobs_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_JOBS_URL", "https://boards-api.greenhouse.io"),
				Description: "The URL for Greenhouse Job Boards API.",
			},
			"on_behalf_of": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GREENHOUSE_ON_BEHALF_OF", nil),
				Description: "This is the user on whose behalf all actions will be audited.",
			},
			"retry_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"retry_wait": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"retry_max_wait": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"greenhouse_application":           resourceGreenhouseApplication(),
			"greenhouse_approval":              resourceGreenhouseApproval(),
			"greenhouse_candidate":             resourceGreenhouseCandidate(),
			"greenhouse_custom_field":          resourceGreenhouseCustomField(),
			"greenhouse_department":            resourceGreenhouseDepartment(),
			"greenhouse_future_job_permission": resourceGreenhouseFutureJobPermission(),
			"greenhouse_hiring_team":           resourceGreenhouseHiringTeam(),
			"greenhouse_job":                   resourceGreenhouseJob(),
			"greenhouse_job_opening":           resourceGreenhouseJobOpening(),
			"greenhouse_job_permission":        resourceGreenhouseJobPermission(),
			"greenhouse_job_post":              resourceGreenhouseJobPost(),
			"greenhouse_offer":                 resourceGreenhouseOffer(),
			"greenhouse_office":                resourceGreenhouseOffice(),
			"greenhouse_scheduled_interview":   resourceGreenhouseScheduledInterview(),
			"greenhouse_tag":                   resourceGreenhouseCandidateTag(),
			"greenhouse_user":                  resourceGreenhouseUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"greenhouse_activity_feed":             dataSourceGreenhouseActivityFeed(),
			"greenhouse_applications":              dataSourceGreenhouseApplications(),
			"greenhouse_approvals":                 dataSourceGreenhouseApprovals(),
			"greenhouse_candidates":                dataSourceGreenhouseCandidates(),
			"greenhouse_close_reason":              dataSourceGreenhouseCloseReason(),
			"greenhouse_close_reasons":             dataSourceGreenhouseCloseReasons(),
			"greenhouse_custom_field_options":      dataSourceGreenhouseCustomFieldOptions(),
			"greenhouse_custom_fields":             dataSourceGreenhouseCustomFields(),
			"greenhouse_demographic_answer":        dataSourceGreenhouseDemographicAnswer(),
			"greenhouse_demographic_answer_option": dataSourceGreenhouseDemographicAnswerOption(),
			"greenhouse_demographic_question":      dataSourceGreenhouseDemographicQuestion(),
			"greenhouse_demographic_questions":     dataSourceGreenhouseDemographicQuestions(),
			"greenhouse_demographic_question_set":  dataSourceGreenhouseDemographicQuestionSet(),
			"greenhouse_department":                dataSourceGreenhouseDepartment(),
			"greenhouse_departments":               dataSourceGreenhouseDepartments(),
			"greenhouse_education_degree":          dataSourceGreenhouseEducationDegree(),
			"greenhouse_education_discipline":      dataSourceGreenhouseEducationDiscipline(),
			"greenhouse_education_school":          dataSourceGreenhouseEducationSchool(),
			"greenhouse_eeoc":                      dataSourceGreenhouseEEOC(),
			"greenhouse_email_template":            dataSourceGreenhouseEmailTemplate(),
			"greenhouse_future_job_permissions":    dataSourceGreenhouseFutureJobPermissions(),
			"greenhouse_job":                       dataSourceGreenhouseJob(),
			"greenhouse_job_openings":              dataSourceGreenhouseJobOpenings(),
			"greenhouse_job_permissions":           dataSourceGreenhouseJobPermissions(),
			"greenhouse_job_posts":                 dataSourceGreenhouseJobPosts(),
			"greenhouse_job_stage":                 dataSourceGreenhouseJobStage(),
			"greenhouse_jobs":                      dataSourceGreenhouseJobs(),
			"greenhouse_offer":                     dataSourceGreenhouseOffer(),
			"greenhouse_offers":                    dataSourceGreenhouseOffers(),
			"greenhouse_office":                    dataSourceGreenhouseOffice(),
			"greenhouse_offices":                   dataSourceGreenhouseOffices(),
			"greenhouse_prospect_pool":             dataSourceGreenhouseProspectPool(),
			"greenhouse_rejection_reason":          dataSourceGreenhouseRejectionReason(),
			"greenhouse_scheduled_interviews":      dataSourceGreenhouseScheduledInterviews(),
			"greenhouse_scorecard":                 dataSourceGreenhouseScorecard(),
			"greenhouse_source":                    dataSourceGreenhouseSource(),
			"greenhouse_sources":                   dataSourceGreenhouseSources(),
			"greenhouse_tags":                      dataSourceGreenhouseCandidateTags(),
			"greenhouse_tracking_link":             dataSourceGreenhouseTrackingLink(),
			"greenhouse_user":                      dataSourceGreenhouseUser(),
			"greenhouse_user_role":                 dataSourceGreenhouseUserRole(),
			"greenhouse_users":                     dataSourceGreenhouseUsers(),
		},
	}
	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		ctx = context.TODO()
		harvest_url := d.Get("harvest_url").(string)
		harvest_token := d.Get("harvest_token").(string)
		on_behalf_of, err := strconv.Atoi(d.Get("on_behalf_of").(string))
		if err != nil {
			return nil, err
		}
		client := greenhouse.Client{
			BaseUrl:    harvest_url,
			Token:      harvest_token,
			OnBehalfOf: on_behalf_of,
		}
		if v, ok := d.Get("retry_count").(int); ok && v != 0 {
			client.RetryCount = v
		} else {
			tflog.Warn(ctx, "Not setting retry count.")
		}
		if v, ok := d.Get("retry_wait").(int64); ok && v != 0 {
			client.RetryWait = v
		} else {
			tflog.Warn(ctx, "Not setting retry wait.")
		}
		if v, ok := d.Get("retry_max_wait").(int64); ok && v != 0 {
			client.RetryMaxWait = v
		} else {
			tflog.Warn(ctx, "Not setting retry max wait.")
		}
		err = client.BuildResty()
		if err != nil {
			return nil, err
		}
		return &client, nil
	}
}

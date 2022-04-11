package greenhouse

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceGreenhouseHiringTeam() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceGreenhouseHiringTeamCreate,
		ReadContext:   resourceGreenhouseHiringTeamRead,
		UpdateContext: resourceGreenhouseHiringTeamUpdate,
		DeleteContext: resourceGreenhouseHiringTeamDelete,
		Exists:        resourceGreenhouseHiringTeamExists,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, client interface{}) ([]*schema.ResourceData, error) {
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: schemaGreenhouseHiringTeam(),
	}
}

func resourceGreenhouseHiringTeamExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return false, err
	}
	return greenhouse.Exists(meta.(*greenhouse.Client), context.TODO(), fmt.Sprintf("v1/jobs/%d/hiring_team", id))
}

func resourceGreenhouseHiringTeamCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceGreenhouseHiringTeamCreate")
	var obj map[string][]greenhouse.HiringMember
	var err error
	jobId := d.Get("job_id").(int)
	if v, ok := d.Get("teams").([]interface{}); ok && len(v) > 0 {
		obj, err = transformHiringTeam(ctx, &v)
		if err != nil {
			return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
		}
	}
	diagErr := logJson(ctx, "resourceGreenhouseHiringTeamCreate", obj)
	if diagErr != nil {
		return diagErr
	}
	err = greenhouse.UpdateJobHiringTeam(meta.(*greenhouse.Client), ctx, jobId, &obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.SetId(strconv.Itoa(jobId))
	return resourceGreenhouseHiringTeamRead(ctx, d, meta)
}

func resourceGreenhouseHiringTeamRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Trace(ctx, "Getting hiring team.")
	obj, err := greenhouse.GetJobHiringTeam(meta.(*greenhouse.Client), ctx, id)
	tflog.Trace(ctx, "Got hiring team.")
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	d.Set("teams", flattenHiringSubteams(ctx, obj))
	return nil
}

func resourceGreenhouseHiringTeamUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceGreenhouseHiringTeamUpdate")
	return resourceGreenhouseHiringTeamCreate(ctx, d, meta)
}

func resourceGreenhouseHiringTeamDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return diag.Diagnostics{{Severity: diag.Error, Summary: "Delete is not supported for hiring_team."}}
}

func transformHiringTeam(ctx context.Context, hiringTeam *[]interface{}) (map[string][]greenhouse.HiringMember, error) {
	update := make(map[string][]greenhouse.HiringMember)
	for _, team := range *hiringTeam {
		teamItem := team.(map[string]interface{})
		teamName := teamItem["name"].(string)
		members := teamItem["members"].([]interface{})
		if len(members) > 0 {
			update[teamName] = make([]greenhouse.HiringMember, len(members), len(members))
			for j, member := range members {
				var obj greenhouse.HiringMember
				marshaled, err := json.Marshal(member)
				if err != nil {
					return nil, err
				}
				err = json.Unmarshal(marshaled, &obj)
				if err != nil {
					return nil, err
				}
				update[teamName][j] = obj
			}
		} else {
			update[teamName] = make([]greenhouse.HiringMember, 0)
		}
	}
	tflog.Trace(ctx, "Updating hiring team", "updateObj", fmt.Sprintf("%+v", update))
	return update, nil
}

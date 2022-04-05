package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"reflect"
	"testing"
)

var (
	testGreenhouseUserObj greenhouse.User
	testGreenhouseUserInt map[string]interface{}
)

func init() {
	ctx = context.TODO()
	testGreenhouseUserObj = greenhouse.User{
		CreatedAt:          StringPtr("01-01-2020"),
		Disabled:           BoolPtr(false),
		Emails:             []string{"a@test.com", "b@test.com"},
		EmployeeId:         StringPtr("test123"),
		FirstName:          StringPtr("test"),
		LastName:           StringPtr("user"),
		LinkedCandidateIds: []int{123},
		Name:               StringPtr("test user"),
		PrimaryEmail:       StringPtr("a@test.com"),
		SiteAdmin:          BoolPtr(false),
		UpdatedAt:          StringPtr("01-02-2020"),
	}
	testGreenhouseUserInt = map[string]interface{}{
		"created_at":            StringPtr("01-01-2020"),
		"disabled":              BoolPtr(false),
		"emails":                []string{"a@test.com", "b@test.com"},
		"employee_id":           StringPtr("test123"),
		"first_name":            StringPtr("test"),
		"last_name":             StringPtr("user"),
		"linked_candidate_ids":  []int{123},
		"name":                  StringPtr("test user"),
		"primary_email_address": StringPtr("a@test.com"),
		"site_admin":            BoolPtr(false),
		"updated_at":            StringPtr("01-02-2020"),
	}
}

func TestFlattenUser(t *testing.T) {
	cases := []struct {
		Input    greenhouse.User
		Expected map[string]interface{}
	}{
		{
			testGreenhouseUserObj,
			testGreenhouseUserInt,
		},
	}
	for _, c := range cases {
		output := flattenUser(ctx, &c.Input)
		if !reflect.DeepEqual(output, c.Expected) {
			t.Fatalf("Failed to flatten. Expected: %+v\nGot: %+v\n", c.Expected, output)
		}
	}
}

func TestInflateUser(t *testing.T) {
	cases := []struct {
		Input    map[string]interface{}
		Expected greenhouse.User
	}{
		{
			testGreenhouseUserInt,
			testGreenhouseUserObj,
		},
	}
	for _, c := range cases {
		output, err := inflateUser(ctx, &c.Input)
		if err != nil {
			t.Fatalf("Error occurred during inflation: %s", err[0].Summary)
		}
		if !reflect.DeepEqual(output, c.Expected) {
			t.Fatalf("Failed to inflate. Expected: %+v\nGot: %+v\n", c.Expected, output)
		}
	}
}

package greenhouse

import (
	"context"
	"github.com/carnegierobotics/greenhouse-client-go/greenhouse"
	"reflect"
	"testing"
)

var (
	ctx                             context.Context
	testGreenhouseUserPermissionObj []greenhouse.UserPermission
	testGreenhouseUserPermissionInt []interface{}
)

func init() {
	ctx = context.TODO()
	testGreenhouseUserPermissionObj = []greenhouse.UserPermission{
		{
			JobId:      123,
			UserRoleId: 789,
		},
	}
	testGreenhouseUserPermissionInt = []interface{}{
		map[string]interface{}{
			"job_id":       123,
			"user_role_id": 789,
		},
	}
}

func TestFlattenUserPermissions(t *testing.T) {
	cases := []struct {
		Input    []greenhouse.UserPermission
		Expected []interface{}
	}{
		{
			testGreenhouseUserPermissionObj,
			testGreenhouseUserPermissionInt,
		},
	}
	for _, c := range cases {
		output := flattenUserPermissions(ctx, &c.Input)
		if !reflect.DeepEqual(output, c.Expected) {
			t.Fatalf("Failed to flatten. Expected: %+v\nGot: %+v\n", c.Expected, c.Input)
		}
	}
}

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
	"reflect"
	"testing"
)

var (
	testGreenhouseUserPermissionObj []greenhouse.UserPermission
	testGreenhouseUserPermissionInt []interface{}
)

func init() {
	ctx = context.TODO()
	testGreenhouseUserPermissionObj = []greenhouse.UserPermission{
		{
			JobId:      IntPtr(123),
			UserRoleId: IntPtr(789),
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
			t.Fatalf("Failed to flatten. Expected: %+v\nGot: %+v\n", c.Expected, output)
		}
	}
}

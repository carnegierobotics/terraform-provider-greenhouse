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
	testGreenhouseDepartmentObj []greenhouse.Department
	testGreenhouseDepartmentInt []interface{}
)

func init() {
	ctx = context.TODO()
	testGreenhouseDepartmentObj = []greenhouse.Department{
		{
			ChildDepartmentExternalIds: []string{"idA", "idB"},
			ChildIds:                   []int{123, 456},
			ExternalId:                 StringPtr("test"),
			Name:                       StringPtr("test"),
			ParentDepartmentExternalId: StringPtr("parentExtId"),
			ParentId:                   IntPtr(1),
		},
	}
	testGreenhouseDepartmentInt = []interface{}{
		map[string]interface{}{
			"child_department_external_ids": []string{"idA", "idB"},
			"child_ids":                     []int{123, 456},
			//"external_id":                   "test",
			"name":                          "test",
			"parent_department_external_id": "parentExtId",
			"parent_id":                     1,
		},
	}
}

func TestFlattenDepartments(t *testing.T) {
	cases := []struct {
		Input    []greenhouse.Department
		Expected []interface{}
	}{
		{
			testGreenhouseDepartmentObj,
			testGreenhouseDepartmentInt,
		},
	}
	for _, c := range cases {
		output := flattenDepartments(ctx, &c.Input)
		if !reflect.DeepEqual(output, c.Expected) {
			t.Fatalf("Failed to flatten. Expected: %+v\nGot: %+v\n", c.Expected, output)
		}
	}
}

func TestInflateDepartments(t *testing.T) {
	cases := []struct {
		Input    []interface{}
		Expected []greenhouse.Department
	}{
		{
			testGreenhouseDepartmentInt,
			testGreenhouseDepartmentObj,
		},
	}
	for _, c := range cases {
		output, err := inflateDepartments(ctx, &c.Input)
		if err != nil {
			t.Fatalf("Error occurred during inflation: %s", err[0].Summary)
		}
		if !reflect.DeepEqual(output, c.Expected) {
			t.Fatalf("Failed to inflate. Expected: %+v\nGot: %+v\n", c.Expected, output)
		}
	}
}

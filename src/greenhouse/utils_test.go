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
	"reflect"
	"testing"
)

var (
	testGreenhouseEducationOld []interface{}
	testGreenhouseEducationNew []interface{}
	testGreenhouseEducationAdd []interface{}
	testGreenhouseEducationDel []interface{}
)

func init() {
	ctx = context.TODO()
	testGreenhouseEducationOld = []interface{}{
		map[string]interface{}{
			"id": 123,
		},
		map[string]interface{}{
			"id": 234,
		},
	}
	testGreenhouseEducationNew = []interface{}{
		map[string]interface{}{
			"id": 234,
		},
		map[string]interface{}{
			"id": 345,
		},
	}
	testGreenhouseEducationAdd = []interface{}{
		map[string]interface{}{
			"id": 345,
		},
	}
	testGreenhouseEducationDel = []interface{}{
		map[string]interface{}{
			"id": 123,
		},
	}
}

func TestFindAddDelete(t *testing.T) {
	cases := []struct {
		InputOld    []interface{}
		InputNew    []interface{}
		ExpectedAdd []interface{}
		ExpectedDel []interface{}
	}{
		{
			testGreenhouseEducationOld,
			testGreenhouseEducationNew,
			testGreenhouseEducationAdd,
			testGreenhouseEducationDel,
		},
	}
	for _, c := range cases {
		add, del, err := findAddDelete(ctx, c.InputOld, c.InputNew)
		if err != nil {
			t.Fatalf("Error: %s.", err[0].Summary)
		}
		if !reflect.DeepEqual(*add, c.ExpectedAdd) {
			t.Fatalf("Failed to find add. Expected: %+v\nGot: %+v\n", c.ExpectedAdd, *add)
		}
		if !reflect.DeepEqual(*del, c.ExpectedDel) {
			t.Fatalf("Failed to find del. Expected: %+v\nGot: %+v\n", c.ExpectedDel, *del)
		}
	}
}

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
			"external_id":                   "test",
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

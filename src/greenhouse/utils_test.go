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

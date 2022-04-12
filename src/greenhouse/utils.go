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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func mapAItoMapAA(ctx context.Context, mapAI map[string]interface{}) *map[string]string {
	mapAA := make(map[string]string)
	for k, v := range mapAI {
		mapAA[k] = v.(string)
	}
	return &mapAA
}

func Bool(ptr *bool) bool {
	if ptr != nil {
		return *ptr
	} else {
		return false
	}
}

func BoolPtr(v bool) *bool {
	return &v
}

func emptyList() []interface{} {
	return make([]interface{}, 0)
}

func findAddDelete(ctx context.Context, o interface{}, n interface{}) (*[]interface{}, *[]interface{}, diag.Diagnostics) {
	v, ok1 := o.([]interface{})
	w, ok2 := n.([]interface{})
	if !ok1 || !ok2 {
		return nil, nil, diag.Diagnostics{{Severity: diag.Error, Summary: "Failed to convert to []interface{}"}}
	}
	add := make([]interface{}, 0)
	del := make([]interface{}, 0)
	for _, i1 := range v {
		match := false
		for _, i2 := range w {
			if i1.(map[string]interface{})["id"] == i2.(map[string]interface{})["id"] {
				match = true
				break
			}
		}
		if !match {
			del = append(del, i1)
		}
	}
	for _, i1 := range w {
		match := false
		for _, i2 := range v {
			if i1.(map[string]interface{})["id"] == i2.(map[string]interface{})["id"] {
				match = true
				break
			}
		}
		if !match {
			add = append(add, i1)
		}
	}
	return &add, &del, nil
}

type readFunc func(context.Context, *schema.ResourceData, interface{}) diag.Diagnostics

func importByRead(ctx context.Context, d *schema.ResourceData, meta interface{}, fn readFunc) ([]*schema.ResourceData, error) {
	err := fn(ctx, d, meta)
	if err != nil {
		return nil, errors.New(err[0].Summary)
	}
	return []*schema.ResourceData{d}, nil
}

func Int(ptr *int) int {
	if ptr != nil {
		return *ptr
	} else {
		return 0
	}
}

func IntPtr(v int) *int {
	return &v
}

func logJson(ctx context.Context, vertex string, obj interface{}) diag.Diagnostics {
	jsonBody, err := json.Marshal(obj)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Trace(ctx, vertex, fmt.Sprintf("JSON will be: %s", string(jsonBody)))
	return nil
}

func String(ptr *string) string {
	if ptr != nil {
		return *ptr
	} else {
		return ""
	}
}

func StringPtr(v string) *string {
	return &v
}

func sliceItoSliceA(sliceI *[]interface{}) *[]string {
	sliceA := make([]string, len(*sliceI), len(*sliceI))
	for i, item := range *sliceI {
		sliceA[i] = item.(string)
	}
	return &sliceA
}

func sliceItoSliceD(sliceI *[]interface{}) *[]int {
	sliceD := make([]int, len(*sliceI), len(*sliceI))
	for i, item := range *sliceI {
		sliceD[i] = item.(int)
	}
	return &sliceD
}

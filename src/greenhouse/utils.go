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

func convertType(ctx context.Context, source interface{}, target interface{}) diag.Diagnostics {
	tflog.Debug(ctx, fmt.Sprintf("Converting source: %+v", source))
	jsonBody, err := json.Marshal(source)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Debug(ctx, fmt.Sprintf("Finished marshal: %s", string(jsonBody)))
	err = json.Unmarshal(jsonBody, &target)
	if err != nil {
		return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
	}
	tflog.Debug(ctx, fmt.Sprintf("Finished conversion: %+v", target))
	return nil
}

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

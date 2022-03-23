package greenhouse

import (
  "context"
  "encoding/json"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func ConvertSliceInterfaceInt(slice []interface{}) []int {
	newslice := make([]int, len(slice), len(slice))
	for i := range slice {
		newslice[i] = slice[i].(int)
	}
	return newslice
}

func ConvertSliceInterfaceString(slice []interface{}) *[]string {
  newslice := make([]string, len(slice), len(slice))
  for i := range slice {
    newslice[i] = slice[i].(string)
  }
  return &newslice
}

func convertType(ctx context.Context, source interface{}, target interface{}) diag.Diagnostics {
  jsonBody, err := json.Marshal(source)
  if err != nil {
    return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
  }
  err = json.Unmarshal(jsonBody, &target)
  if err != nil {
    return diag.Diagnostics{{Severity: diag.Error, Summary: err.Error()}}
  }
  return nil
}

package greenhouse

import()

func ConvertSliceInterfaceInt(slice []interface{}) []int {
  newslice := make([]int, len(slice), len(slice))
  for i := range slice {
    newslice[i] = slice[i].(int)
  }
  return newslice
}

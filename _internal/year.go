package internal

import (
  "fmt"
  "github.com/spf13/cast"
)

func yearOrDefault(year interface{}, defaultYear int) string {
  if yearValue, ok := year.(string); ok {
    return yearValue
  } else if yearValue, ok := year.(int); ok {
    return fmt.Sprintf("%d", yearValue)
  }
  return fmt.Sprintf("%d", defaultYear)
}
package infrastructure

import (
	"fmt"
	"reflect"
	"strconv"
)

//ConvertNumberToString ..
func ConvertNumberToString(arg interface{}) string {
	switch reflect.ValueOf(arg).Kind() {
	case reflect.Int:
		return strconv.Itoa(arg.(int))
	case reflect.Float32:
		return fmt.Sprintf("%v", arg.(float32))
	default:
		return ""
	}
}

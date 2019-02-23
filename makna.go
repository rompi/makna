package makna

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToString convert data types in golang to string
func ToString(val interface{}) string {
	return fmt.Sprint(val)
}

// ToInt32E convert data types in golang to int32 with error
func ToInt32E(val interface{}) (int32, error) {
	switch valType := reflect.TypeOf(val).String(); valType {
	case "string":
		val, err := strconv.ParseInt(fmt.Sprint(val), 10, 32)
		return int32(val), err
	default:
		return 0, fmt.Errorf(CAST_ERROR, valType)
	}
}

// ToInt32 convert data types in golang to int32 (without returning the error)
func ToInt32(val interface{}) int32 {
	newVal, _ := ToInt32E(val)
	return newVal
}

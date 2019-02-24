package makna

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToString convert data types in golang to string
func ToString(val interface{}) string {
	switch valType := reflect.TypeOf(val).String(); valType {
	case "float64":
		return fmt.Sprintf("%6.2f", val)
	default:
		return fmt.Sprintf("%+v", val)
	}
}

// ToInt32 convert data types in golang to int32 (without returning the error)
func ToInt32(val interface{}) int32 {
	newVal, _ := ToInt32E(val)
	return newVal
}

// ToInt32E convert data types in golang to int32 with error
func ToInt32E(val interface{}) (int32, error) {
	newVal, err := strconv.ParseInt(fmt.Sprintf("%+v", val), 10, 32)
	return int32(newVal), err
}

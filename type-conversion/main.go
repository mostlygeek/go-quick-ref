package main

// type conversions are checked at compile time. Thus why all the boilerplate
// code for the ToInt() and ToFloat64 functions.

import (
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(v interface{}) (int, bool) {

	// check if it implements a right type of interface
	// to convert to an int
	switch v.(type) {
	// most common cases first
	case int:
		return v.(int), true
	case float64:
		return int(v.(float64)), true
	case uint:
		return int(v.(uint)), true
	case string:
		converted, err := strconv.Atoi(v.(string))
		if err == nil {
			return converted, true
		} else {
			return 0, false
		}

	case int8:
		return int(v.(int8)), true
	case int16:
		return int(v.(int16)), true
	case int32:
		return int(v.(int32)), true
	case int64:
		return int(v.(int64)), true
	case uint8:
		return int(v.(uint8)), true
	case uint16:
		return int(v.(uint16)), true
	case uint32:
		return int(v.(uint32)), true
	case uint64:
		return int(v.(uint64)), true
	case float32:
		return int(v.(float32)), true
	default:
		return 0, false
	}

}

func ToFloat64(v interface{}) (float64, bool) {

	switch v.(type) {
	case float64:
		return v.(float64), true
	case float32:
		return float64(v.(float32)), true

	case int:
		return float64(v.(int)), true
	case uint:
		return float64(v.(uint)), true

	case string:
		converted, err := strconv.ParseFloat(v.(string), 64)
		if err == nil {
			return converted, true
		} else {
			return 0, false
		}

	case int8:
		return float64(v.(int8)), true
	case int16:
		return float64(v.(int16)), true
	case int32:
		return float64(v.(int32)), true
	case int64:
		return float64(v.(int64)), true
	case uint8:
		return float64(v.(uint8)), true
	case uint16:
		return float64(v.(uint16)), true
	case uint32:
		return float64(v.(uint32)), true
	case uint64:
		return float64(v.(uint64)), true
	default:
		return 0, false
	}
}

func main() {
	vals := []interface{}{
		int(1), int8(1), int16(1), int32(1), int64(1),
		uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
		float64(1.29), float32(1.29),

		// strings to try out
		"hello",
		"1",
		"3.14",

		true,
		false,
		[]int{1, 2, 3}, map[string]int{"k": 1},
	}
	for _, v := range vals {
		if converted, ok := ToInt(v); ok {
			fmt.Println(v, "is a", reflect.TypeOf(v).Kind(), ",converted to int ==", converted)
		} else {
			fmt.Println("X - the", reflect.TypeOf(v).Kind(), v, "can't be converted to an int")
		}
	}

	fmt.Println("\n------\n")

	for _, v := range vals {
		if converted, ok := ToFloat64(v); ok {
			fmt.Println(v, "is a", reflect.TypeOf(v).Kind(), ",converted to float64 ==", converted)
		} else {
			fmt.Println("X - the", reflect.TypeOf(v).Kind(), v, "can't be converted to an float64")
		}
	}
}

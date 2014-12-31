package main

// demo of a map that can hold any value type, like a JSON object

import (
	"fmt"
	"reflect"
)

// it holds ... anything
type Params map[string]interface{}

func main() {

	m := Params{
		"k1": "a string",
		"k2": 123,
		"k3": true,
		"k4": 9.13,
	}

	for k, v := range m {
		fmt.Printf("%v := %v (%v)\n", k, v, reflect.TypeOf(v))
	}

	// another way of properly figuring out the type
	fmt.Println("Using Type Assertions")
	for k, v := range m {
		var t string

		switch v.(type) {
		case string:
			t = "string"
		case int:
			t = "int"
		case bool:
			t = "bool"
		default:
			t = "unknown"
		}

		fmt.Printf("%v is a %v\n", k, t)
	}

}

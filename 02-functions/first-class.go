package main

// demonstration of first class / more advanced function usage in go

import (
	"fmt"
)

// Declare a new Type
type Validator func(interface{}) bool

// Create a function that uses a closure
func ValidateStringEquals(a string) Validator {
	return func(v interface{}) bool {
		b, ok := v.(string)
		if !ok {
			return false
		}

		return a == b
	}
}

func main() {

	f1 := ValidateStringEquals("foo")
	f2 := ValidateStringEquals("bar")

	// true
	fmt.Printf("f1: %v\n", f1("foo"))
	fmt.Printf("f2: %v\n", f2("bar"))

	// false
	fmt.Printf("f1: %v\n", f1("wrong"))
	fmt.Printf("f2: %v\n", f2("wrong"))

}

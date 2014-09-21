package main

import (
	"fmt"
)

// takes any sort of value for `v` and figures out what
// type it is and creates the appropriate output
func TakesSomething(label string, v interface{}) string {

	switch v.(type) {
	case string:
		return fmt.Sprintf("%s: string", label)
	case int, int32, int64:
		return fmt.Sprintf("%s: int of some kind", label)
	case float32, float64:
		return fmt.Sprintf("%s: float of some kind", label)
	default:
		return fmt.Sprintf("%s: no idea", label)
	}
}

// example of dynamic typing in go
// ref: http://blog.denevell.org/golang-interface-type-assertions-switch.html
func main() {

	fmt.Println(TakesSomething("1", 1235))
	fmt.Println(TakesSomething("2", "hello world"))
	fmt.Println(TakesSomething("3", 12.31))
	fmt.Println(TakesSomething("4", true))

	// demo of handling a type as something
	var myVar interface{} = 1234

	// this will panic in go: aString := myVar.(string)
	_, ok := myVar.(string)
	if !ok {
		fmt.Println("5: could not convert that int to a string")
	}

	aNum, ok := myVar.(int)
	if ok {
		fmt.Printf("6: myVar=%d\n", aNum)
	}

	// or terse way of dealing with conditional and type conversion
	if aNum, ok := myVar.(int); ok {
		fmt.Printf("7: terse conversion ok (%d)\n", aNum)
	}

}

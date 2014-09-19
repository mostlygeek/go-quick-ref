package main

import (
    "fmt"
)

func addOne(x int) int /* return type */ {
    return x + 1
}

func oneOff(x int) (int, int) /* multiple return types */ {
    return x - 1, x + 1
}

// named result params: http://golang.org/doc/effective_go.html#named-results
func twoOff(x int) (a, b int) {
    a = x - 2       // you can assign values to them!
    b = x + 2
    return
}

func stringAndNumber() (s string, i int) {
    s = "hello"
    i = 5
    return
}

func main() {
    fmt.Printf("addOne(1) = %v\n", addOne(1))

    a,b := oneOff(1)
    fmt.Printf("oneOff(1) = %v, %v\n", a, b)

    c,d := twoOff(2)
    fmt.Printf("twoOff(2) = %v, %v\n", c, d)

    e,f := stringAndNumber()
    fmt.Printf("stringAndNumber() = %v, %v\n", e, f)
}

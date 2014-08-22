package main

import (
    "fmt"
)

func main() {
    // http://www.golang-book.com/8/index.htm

    // p holds the memory address to an integer
    var p *int;

    // hey x is an integer
    x := 10

    // assign address of x to p
    p = &x

    // lets see what they look like..
    fmt.Printf("p  = %v (mem address)\n", p)
    fmt.Printf("*p = %v (actual value)\n", *p) // `*` deferences a pointer

    // assign a new int value at the memory address
    *p = 11

    // x is now equal to 11
    fmt.Printf("x = %v\n", x)

}

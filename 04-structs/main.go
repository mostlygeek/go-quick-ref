package main

import "fmt"

// Use these examples to improve the examples below
// See more: http://golang.org/ref/spec#Struct_types
// Structs / Interfaces: http://www.golang-book.com/9

type MyString string
type InnerStruct struct {
	num    int
	bignum int64
}

type Rectangle struct {
	width, height int
	name          string

	// anonymous fields
	// ref: http://golangtutorials.blogspot.ca/2011/06/anonymous-fields-in-structs-like-object.html
	MyString
	InnerStruct
}

/* r is passed as a value (a copy) */
func (r Rectangle) SetNameAsValue(v string) {
	r.name = v
}

/* this takes a *POINTER* to r  k

In Go, everything is passed by value. Everything.

There are some types (pointers, channels, maps, slices) that have
reference-like properties, but in those cases the relevant data
structure (pointer, channel pointer, map header, slice header) holds a
pointer to an underlying, shared object (pointed-to thing, channel
descriptor, hash table, array); the data structure itself is passed by
value. Always.

Always.
Rob Pike
Src: https://groups.google.com/forum/#!topic/golang-nuts/INedfATw74A

*/
func (r *Rectangle) SetNameAsPointer(v string) {
	r.name = v
}

/*
 What is the difference with a method that uses a pointer as the receiver?
  - the method that uses a pointer can modify values
  - the method that gets a copy cannot modify the original's values

 see: http://stackoverflow.com/a/9958532
*/

func (r *Rectangle) Area() (area int) {
	area = r.width * r.height
	return
}

func main() {

	r1 := Rectangle{10, 10, "my rectangle", "val for MyString", InnerStruct{10, 13}}

	fmt.Println(r1)

	r1.SetNameAsValue("Name1")
	fmt.Println(r1.name) // "my rectangle", not changed!

	r1.SetNameAsPointer("Name2")
	fmt.Println(r1.name) // "Name2"

	r2 := Rectangle{width: 13, height: 9, MyString: "hi", InnerStruct: InnerStruct{num: 11, bignum: 100}}

	fmt.Println(r2)        // {13 9  hi {11 100}}}
	fmt.Println(r2.Area()) // 117

	// another way of creating rectangle with passing
	// of variables w/ correct types
	var v1 MyString = "boom"
	v2 := InnerStruct{11, 1231}
	r3 := Rectangle{13, 9, "name", v1, v2}
	fmt.Println(r3)        // {13 9 name boom {11 1231}}}
	fmt.Println(r3.Area()) // 117
}

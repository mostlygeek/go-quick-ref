package main

import (
	"fmt"
)

type StringMaker interface {
	Make() string
}

type Something1 struct {
	msg string
}

func (s Something1) Make() string {
	return fmt.Sprintf("Something1's msg: %v", s.msg)
}

type Something2 struct {
	num int
}

func (s Something2) Make() string {
	return fmt.Sprintf("Something2's int: %v", s.num)
}

func PrintMessage(s StringMaker) {
	fmt.Println(s.Make())
}

func main() {

	s1 := Something1{"Hello!"}
	s2 := Something2{5}

	PrintMessage(s1)
	PrintMessage(s2)

	// example of using interface to hide specifics with an array
	fmt.Println("\nExample using array:")
	var somethings [2]StringMaker
	somethings[0] = s1
	somethings[1] = s2

	for _, s := range somethings {
		PrintMessage(s)
	}

}

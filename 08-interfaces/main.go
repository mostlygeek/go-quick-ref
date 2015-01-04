package main

import (
	"fmt"
)

type StringMaker interface {
	Make() string
	Append(string) string
}

type Something1 struct {
	msg string
}

func (s Something1) Make() string {
	return fmt.Sprintf("Something1's msg: %v", s.msg)
}

func (s Something1) Append(a string) string {
	return fmt.Sprintf("something1.append(), %v%v", s, a)
}

type Something2 struct {
	num int
}

func (s Something2) Make() string {
	return fmt.Sprintf("Something2's int: %v", s.num)
}
func (s Something2) Append(a string) string {
	return fmt.Sprintf("something2.append(), %v%v", s, a)
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
		fmt.Println(s.Append("hello"))
	}

}

package main

// Demo of Base > Child relationship with
// function from Base being called from a Child

import (
	"fmt"
)

type Toggleable interface {
	Enabled() (bool, string)
}

type Base struct {
	enabled bool
	name    string
}

func (b *Base) Name() string { return b.name }

// implement the Toggleable interface on base
// it will also work on Child!
func (b Base) Enabled() (bool, string) { return b.enabled, "base" }

type Child struct {
	Base
	age int
}

func (c Child) Age() int { return c.age }

type Child2 struct{ Base }

func (c Child2) Enabled() (bool, string) { return c.enabled, "child" }

func IsEnabled(t Toggleable) string {
	if enabled, msg := t.Enabled(); enabled == true {
		return "yes, " + msg
	} else {
		return "no, " + msg
	}

}

func main() {
	b := Base{false, "Base"}
	fmt.Printf("%v is enabled? %v\n", b.Name(), IsEnabled(b))

	//c.Name() works because it includes Base as an
	//anonymous field

	c := Child{Base: Base{true, "Child"}, age: 10}
	fmt.Printf("%v:%v is enabled? %v\n", c.Name(), c.Age(), IsEnabled(c))

	// Child2 has an Enabled() method so that is used instead of the
	// one in Base like with Child1
	c2 := Child2{Base: Base{true, "Child2"}}
	fmt.Printf("Child2 %v is enabled? %v\n", c2.Name(), IsEnabled(c2))

}

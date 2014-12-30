package main

// Demo of Base > Child relationship with
// function from Base being called from a Child

import (
	"fmt"
)

type Toggleable interface {
	Enabled() bool
}

type Base struct {
	enabled bool
	name    string
}

func (b *Base) Name() string { return b.name }

// implement the Toggleable interface on base
// it will also work on Child!
func (b Base) Enabled() bool { return b.enabled }

type Child struct {
	Base
	age int
}

func (c Child) Age() int { return c.age }

func IsEnabled(t Toggleable) string {
	if t.Enabled() {
		return "yes"
	}

	return "no"
}

func main() {
	b := Base{false, "hello"}
	fmt.Printf("%v is enabled? %v\n", b.Name(), IsEnabled(b))

	//c.Name() works because it includes Base as an
	//anonymous field

	c := Child{Base: Base{true, "Child"}, age: 10}
	fmt.Printf("Child %v:%v is enabled? %v\n", c.Name(), c.Age(), IsEnabled(c))
}

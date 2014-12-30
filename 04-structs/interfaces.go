package main

// Demo of Base > Child relationship with
// function from Base being called from a Child

import (
	"fmt"
)

type Base struct {
	name string
}

func (b Base) Name() string {
	return b.name
}

type Child struct {
	Base
	age int
}

func (c Child) Age() int { return c.age }

func main() {
	b := Base{"hello"}
	fmt.Println(b.Name())

	//c.Name() works because it includes Base as an
	//anonymous field

	c := Child{Base: Base{"Child"}, age: 10}
	fmt.Printf("%v:%v\n", c.Name(), c.Age())
}

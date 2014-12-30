package main

// Demo of Base > Child relationship with
// methods from Base being called from a Child
// as well as interactions with an interface

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

func (b *Base) Name() string           { return b.name }
func (b Base) Enabled() (bool, string) { return b.enabled, "[Enabled() from Base]" }

type Child struct{ Base }

// Child2 has its own Enabled() method that hides Base.Enabled()
type Child2 struct{ Base }

func (c Child2) Enabled() (bool, string) { return c.enabled, "[Enabled() from Child2]" }

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

	c := Child{Base{true, "Child"}}
	fmt.Printf("%v is enabled? %v\n", c.Name(), IsEnabled(c))

	// Child2 has an Enabled() method so that is used instead of the
	// one in Base like with Child1
	c2 := Child2{Base{true, "Child2"}}
	fmt.Printf("Child2 %v is enabled? %v\n", c2.Name(), IsEnabled(c2))
}

package main

import (
	"fmt"

	/** Ref:
	 * https://golang.org/ref/spec#Import_declarations
	 */

	/* will be mypackage.Func */
	"github.com/mostlygeek/go-quick/07-packages/mypackage"

	/* alias the package name */
	an_alias "github.com/mostlygeek/go-quick/07-packages/mypackage"

	// import into the local namespace so we can just
	// use Add(...) directly
	. "github.com/mostlygeek/go-quick/07-packages/mypackage"
	. "github.com/mostlygeek/go-quick/07-packages/mypackage/subpackage"
)

func main() {
	fmt.Println(mypackage.Add(1, 1))

	// the alias
	fmt.Println(an_alias.Add(1, 1))

	// local name space imported fns
	fmt.Println(Add(1, 1))
	fmt.Println("Subtract(5, 1) =", Subtract(5, 1))
}

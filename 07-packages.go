package main

import (
    "fmt"

    /** Ref:
     * https://golang.org/ref/spec#Import_declarations
     */

    /* will be package_example.Func */
    "github.com/mostlygeek/go-quick/package_example"

    /* alias the package name */
    an_alias "github.com/mostlygeek/go-quick/package_example"

    // import into the local namespace so we can just
    // use Add(...) directly
    . "github.com/mostlygeek/go-quick/package_example"
)

func main() {
    fmt.Println(package_example.Add(1, 1))

    // the alias
    fmt.Println(an_alias.Add(1, 1))

    // local name space imported fns
    fmt.Println(Add(1, 1))
}

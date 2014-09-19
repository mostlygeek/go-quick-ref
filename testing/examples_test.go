// see: http://golang.org/pkg/testing/
// to see these tests in action, run `go test`
package test_example

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 1) != 2 {
		t.Error("1 + 1 != 2")
	}
}

// examples run and verify the output, which is defined
// after the `// Output:` header
func Example_HelloWorld() {
	HelloWorld()
	// Output:
	// Hello World
}

// run with `go test -bench .` to see this
// the command is `go test -bench <regex>`
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 1)
	}
}

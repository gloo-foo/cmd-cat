package cat_test

import (
	"fmt"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-cat"
)

func ExampleCat_basic() {
	// echo "Hello World\nThis is a test" | cat
	output, _ := testable.Test(command.Cat(), "Hello World\nThis is a test\n")
	fmt.Print(output)
	// Output:
	// Hello World
	// This is a test
}

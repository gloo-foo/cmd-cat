package cat_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-cat"
	"github.com/gloo-foo/testable"
)

func ExampleCat_numberLines() {
	// echo "Line one\nLine two\nLine three" | cat -n
	output, _ := testable.Test(command.Cat(command.CatNumberLines), "Line one\nLine two\nLine three\n")
	fmt.Print(output)
	// Output:
	//      1	Line one
	//      2	Line two
	//      3	Line three
}

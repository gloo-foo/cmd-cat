package cat_test

import (
	"fmt"

	command "github.com/gloo-foo/cmd-cat"
	"github.com/gloo-foo/testable"
)

func ExampleCat_numberNonBlank() {
	// echo "alpha\n\nbeta\n\ngamma" | cat -b
	output, _ := testable.Test(command.Cat(command.CatNumberNonBlank), "alpha\n\nbeta\n\ngamma\n")
	fmt.Print(output)
	// Output:
	//      1	alpha
	//
	//      2	beta
	//
	//      3	gamma
}

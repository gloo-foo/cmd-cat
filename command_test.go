package command_test

import (
	"fmt"
	"testing"

	"github.com/gloo-foo/testable"

	command "github.com/gloo-foo/cmd-cat"
)

func TestCat_Passthrough(t *testing.T) {
	lines, err := testable.TestLines(command.Cat(), "hello\nworld\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"hello", "world"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d: %v", len(lines), len(want), lines)
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestCat_NumberLines(t *testing.T) {
	lines, err := testable.TestLines(command.Cat(command.CatNumberLines), "alpha\n\nbeta\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"     1\talpha", "     2\t", "     3\tbeta"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d: %v", len(lines), len(want), lines)
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestCat_NumberNonBlank(t *testing.T) {
	lines, err := testable.TestLines(command.Cat(command.CatNumberNonBlank), "alpha\n\nbeta\n")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"     1\talpha", "", "     2\tbeta"}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d: %v", len(lines), len(want), lines)
	}
	for i, w := range want {
		if lines[i] != w {
			t.Errorf("line %d: got %q, want %q", i, lines[i], w)
		}
	}
}

func TestCat_EmptyInput(t *testing.T) {
	lines, err := testable.TestLines(command.Cat(), "")
	if err != nil {
		t.Fatal(err)
	}
	if len(lines) != 0 {
		t.Fatalf("got %d lines, want 0: %v", len(lines), lines)
	}
}

func ExampleCat() {
	lines, _ := testable.TestLines(command.Cat(), "hello\nworld\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	// hello
	// world
}

func ExampleCat_numberLines() {
	lines, _ := testable.TestLines(command.Cat(command.CatNumberLines), "hello\nworld\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	// Output:
	//      1	hello
	//      2	world
}

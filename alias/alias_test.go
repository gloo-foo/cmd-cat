package alias_test

import (
	"slices"
	"testing"

	cat "github.com/gloo-foo/cmd-cat/alias"
	"github.com/gloo-foo/testable"
)

// The alias package re-exports the constructor and flag constants under
// unprefixed names. A mis-wired re-export (say, NumberLines bound to the
// disabled constant, or Cat bound to the wrong function) compiles cleanly, so
// only behavior can prove the wiring. Each test exercises one re-export and
// asserts the GNU cat output it must produce.

const numberingInput = "alpha\n\nbeta\n"

func assertLines(t *testing.T, got, want []string) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestAlias_DefaultPassesThrough(t *testing.T) {
	lines, err := testable.TestLines(cat.Cat(), numberingInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"alpha", "", "beta"})
}

func TestAlias_NumberLinesNumbersEveryLine(t *testing.T) {
	// -n numbers all lines, including the blank one.
	lines, err := testable.TestLines(cat.Cat(cat.NumberLines), numberingInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"     1\talpha", "     2\t", "     3\tbeta"})
}

func TestAlias_NumberNonBlankSkipsBlankLines(t *testing.T) {
	// -b numbers only non-blank lines; the blank line is left untouched.
	lines, err := testable.TestLines(cat.Cat(cat.NumberNonBlank), numberingInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"     1\talpha", "", "     2\tbeta"})
}

func TestAlias_DisabledFlagsMatchDefault(t *testing.T) {
	// The No* constants are the disabled forms: they must behave exactly like
	// passing no flag at all.
	lines, err := testable.TestLines(cat.Cat(cat.NoNumberLines, cat.NoNumberNonBlank), numberingInput)
	if err != nil {
		t.Fatal(err)
	}
	assertLines(t, lines, []string{"alpha", "", "beta"})
}

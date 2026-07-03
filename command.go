package command

import (
	"fmt"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// lineNumber is the running count a numbered cat run prefixes to each line.
type lineNumber uint64

// Cat returns a Command that concatenates input to output with optional line
// numbering.
//
// Flags:
//   - CatNumberLines (-n): number all output lines
//   - CatNumberNonBlank (-b): number non-blank lines only (overrides -n)
func Cat(opts ...any) gloo.Command[[]byte, []byte] {
	f := fold(opts)
	return patterns.StatefulMap(func() func([]byte) ([]byte, error) {
		// The factory runs once per Execute, giving each pipeline its own
		// fresh counter, so a Command is safe to reuse.
		var n lineNumber
		return func(line []byte) ([]byte, error) {
			out, next := f.number(n, line)
			n = next
			return out, nil
		}
	})
}

// isNumbering reports whether any numbering flag is on (-n or -b).
func (f flags) isNumbering() bool {
	return bool(f.shouldNumberLines) || bool(f.shouldNumberNonBlank)
}

// number transforms one line against the running counter n, returning the
// rendered line and the counter to carry to the next line. Under -b, blank
// lines are neither numbered nor counted.
func (f flags) number(n lineNumber, line []byte) ([]byte, lineNumber) {
	if bool(f.shouldNumberNonBlank) && len(line) == 0 {
		return line, n
	}
	n++
	if !f.isNumbering() {
		return line, n
	}
	return prefix(n, line), n
}

// prefix renders a GNU-cat line number ("%6d\t") in front of line.
func prefix(n lineNumber, line []byte) []byte {
	return append(fmt.Appendf(nil, "%6d\t", n), line...)
}

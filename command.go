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
	f := gloo.NewParameters[gloo.File, flags](opts...).Flags
	return patterns.StatefulMap(func() func([]byte) ([]byte, error) {
		run := &numbering{
			enabled:   bool(f.numberLines) || bool(f.numberNonBlank),
			skipBlank: bool(f.numberNonBlank),
		}
		return run.next
	})
}

// numbering holds the per-Execute state of a cat run. The factory in Cat makes a
// fresh value for every pipeline execution, so a Command is safe to reuse.
//
//   - enabled:   prefix a line number (set by -n or -b).
//   - skipBlank: blank lines are neither numbered nor counted (set by -b).
type numbering struct {
	enabled   bool
	skipBlank bool
	count     lineNumber
}

// next transforms one input line. Pointer receiver: it advances the per-run
// counter, which is the method's contract (the framework's StatefulMap state).
func (n *numbering) next(line []byte) ([]byte, error) {
	if n.skipBlank && len(line) == 0 {
		return line, nil
	}
	n.count++
	if !n.enabled {
		return line, nil
	}
	return prefix(n.count, line), nil
}

// prefix renders a GNU-cat line number ("%6d\t") in front of line.
func prefix(n lineNumber, line []byte) []byte {
	return append(fmt.Appendf(nil, "%6d\t", n), line...)
}

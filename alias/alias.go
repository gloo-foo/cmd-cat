// Package alias provides unprefixed names for cat command flags.
//
//	import cat "github.com/gloo-foo/cmd-cat/alias"
//	cat.Cat(cat.NumberLines)
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-cat"
)

// Cat concatenates input to output with optional line numbering; see the
// command package for the flag set.
func Cat(opts ...any) gloo.Command[[]byte, []byte] { return command.Cat(opts...) }

// -n flag: number all lines
const NumberLines = command.CatNumberLines

// default: don't number lines
const NoNumberLines = command.CatNoNumberLines

// -b flag: number non-blank lines only
const NumberNonBlank = command.CatNumberNonBlank

// default: don't number non-blank lines
const NoNumberNonBlank = command.CatNoNumberNonBlank

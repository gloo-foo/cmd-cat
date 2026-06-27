// Package alias provides unprefixed type aliases for cat command flags.
//
//	import cat "github.com/gloo-foo/cmd-cat/alias"
//	cat.Cat(cat.NumberLines)
package alias

import command "github.com/gloo-foo/cmd-cat"

// Cat re-exports the constructor.
var Cat = command.Cat

// -n flag: number all lines
const NumberLines = command.CatNumberLines

// default: don't number lines
const NoNumberLines = command.CatNoNumberLines

// -b flag: number non-blank lines only
const NumberNonBlank = command.CatNumberNonBlank

// default: don't number non-blank lines
const NoNumberNonBlank = command.CatNoNumberNonBlank

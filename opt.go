package command

// catNumberLinesFlag enables numbering all output lines (-n).
type catNumberLinesFlag bool

const (
	CatNumberLines   catNumberLinesFlag = true
	CatNoNumberLines catNumberLinesFlag = false
)

// catNumberNonBlankFlag enables numbering non-blank lines only (-b).
type catNumberNonBlankFlag bool

const (
	CatNumberNonBlank   catNumberNonBlankFlag = true
	CatNoNumberNonBlank catNumberNonBlankFlag = false
)

// flags is the option set folded from a Cat call's option values.
type flags struct {
	shouldNumberLines    catNumberLinesFlag
	shouldNumberNonBlank catNumberNonBlankFlag
}

// with folds one option value into the flag set. Values of any other type are
// ignored: cat's line transform takes no positional arguments.
func (f flags) with(o any) flags {
	switch v := o.(type) {
	case catNumberLinesFlag:
		f.shouldNumberLines = v
	case catNumberNonBlankFlag:
		f.shouldNumberNonBlank = v
	}
	return f
}

// fold collapses the Cat option values into the flag set.
func fold(opts []any) flags {
	var f flags
	for _, o := range opts {
		f = f.with(o)
	}
	return f
}

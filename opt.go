package command

// catNumberLinesFlag enables numbering all output lines (-n).
type catNumberLinesFlag bool

const (
	CatNumberLines   catNumberLinesFlag = true
	CatNoNumberLines catNumberLinesFlag = false
)

func (f catNumberLinesFlag) Configure(flags *flags) { flags.numberLines = f }

// catNumberNonBlankFlag enables numbering non-blank lines only (-b).
type catNumberNonBlankFlag bool

const (
	CatNumberNonBlank   catNumberNonBlankFlag = true
	CatNoNumberNonBlank catNumberNonBlankFlag = false
)

func (f catNumberNonBlankFlag) Configure(flags *flags) { flags.numberNonBlank = f }

type flags struct {
	numberLines    catNumberLinesFlag
	numberNonBlank catNumberNonBlankFlag
}

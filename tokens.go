package abc

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS
	NL

	// characters
	SEMICOLON     // :
	BAR_SEPARATOR // |
	PERCENT       // %
	LC            // (
	RC            // )
	LS            // ]
	RS            // [

	TEXT // string
)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isNewline(ch rune) bool {
	return ch == '\n'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

var eof = rune(0)

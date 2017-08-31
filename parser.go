package abc

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

var Standard21 = map[string]*InformationField{
	"A": {"A", "area", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},                  //	A:Donegal, A:Bampton (deprecated)
	"B": {"B", "book", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},                  //	B:O'Neills
	"C": {"C", "composer", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},              //	C:Robert Jones, C:Trad.
	"D": {"D", "discography", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},           //	D:Chieftains IV
	"F": {"F", "file", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},                  //	F:http://a.b.c/file.abc
	"G": {"G", "group", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},                 //	G:flute
	"H": {"H", "history", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},               //	H:The story behind this tune …
	"I": {"I", "instruction", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},      // I:papersize A4, I:newpage
	"K": {"K", "key", NOWHERE, LAST, ANYWHERE, true, InstructionValue},                   // K:G, K:Dm, K:AMix
	"L": {"L", "unit note length", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}, // L:1/4, L:1/8
	"M": {"M", "meter", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},            // M:3/4, M:4/4
	"m": {"m", "macro", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},            // m: ~G2 = {A},G{F},G
	"N": {"N", "notes", ANYWHERE, ANYWHERE, ANYWHERE, true, StringValue},                 // N:see also O'Neills - 234
	"O": {"O", "origin", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},                // O:UK; Yorkshire; Bradford
	"P": {"P", "parts", NOWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},             // P:A, P:ABAC, P:(A2B)3
	"Q": {"Q", "tempo", NOWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},             // Q:"allegro" 1/4=120
	"R": {"R", "rhythm", ANYWHERE, ANYWHERE, ANYWHERE, true, StringValue},                // R:R, R:reel
	"r": {"r", "remark", ANYWHERE, ANYWHERE, ANYWHERE, true, StringValue},                // r:I love abc
	"S": {"S", "source", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},                // S:collected in Brittany
	"s": {"s", "symbol line", NOWHERE, NOWHERE, ANYWHERE, false, InstructionValue},       // s: !pp! ** !f!
	"T": {"T", "tune title", NOWHERE, SECOND, ANYWHERE, false, StringValue},              // T:Paddy O'Rafferty
	"U": {"U", "user defined", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},     // U: T = !trill!
	"V": {"V", "voice", NOWHERE, ANYWHERE, ANYWHERE, true, InstructionValue},             // V:4 clef=bass
	"W": {"W", "words", NOWHERE, ANYWHERE, ANYWHERE, false, StringValue},                 // W:lyrics printed after the end of the tune
	"w": {"w", "words", NOWHERE, NOWHERE, ANYWHERE, false, StringValue},                  // w:lyrics printed aligned with the notes of a tune
	"X": {"X", "reference number", NOWHERE, FIRST, NOWHERE, false, InstructionValue},     // X:1, X:2
	"Z": {"Z", "transcription", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue},         // John, j.s@gmail.com
}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
	isHeader bool
	isBody   bool
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{
		s:        NewScanner(r),
		isHeader: false,
		isBody:   false,
	}
}

func (p *Parser) ParseHeader() ([]*Field, error) {
	p.isHeader = true
	fields := []*Field{}
	first, err := p.ParseField()
	if err != nil {
		return fields, err
	} else if first.Shadow.Key != "X" {
		return fields, fmt.Errorf("X must be first declaration")
	}

	fields = append(fields, first)

	second, err := p.ParseField()
	if err != nil {
		return fields, err
	} else if second.Shadow.Key != "T" {
		return fields, fmt.Errorf("T must be second declaration")
	}

	fields = append(fields, second)

LOOP:
	for {
		field, err := p.ParseField()
		if err != nil {
			return fields, err
		}
		if field != nil {
			fields = append(fields, field)
			if field.Shadow.Key == "K" {
				p.isHeader = false
				break LOOP
			}
		} else {
			break LOOP
		}
	}
	if fields[len(fields)-1].Shadow.Key != "K" {
		return fields, fmt.Errorf("K must be last declaration")
	}
	return fields, nil
}

// Parse parses a InformationFieldValue
func (p *Parser) ParseField() (*Field, error) {
	p.isHeader = true
	field := &Field{}
	if tok, lit := p.scanIgnoreWhitespace(); tok != TEXT {
		if tok == EOF {
			return nil, nil
		}
		return nil, fmt.Errorf("found %q, expected TEXT", lit)
	} else {
		field.Shadow = Standard21[lit]
		//fmt.Printf("TOK: %d LIT: %s\n", tok, lit)
	}
	if tok, lit := p.scan(); tok != SEMICOLON {
		return nil, fmt.Errorf("found %q, expected SEMICOLON", lit)
	} else {
		//fmt.Printf("TOK: %d LIT: %s\n", tok, lit)
	}

	var buffer bytes.Buffer
	for {
		tok, lit := p.scanIgnoreWhitespace()
		//fmt.Printf("TOK: %d LIT: %s\n", tok, lit)
		if tok == NL || tok == EOF {
			break
		}
		if tok != TEXT {
			return field, fmt.Errorf("found %q, expected TEXT", lit)
		}

		buffer.WriteString(lit + " ")
	}

	field.Value = strings.TrimSuffix(buffer.String(), " ")

	return field, nil
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

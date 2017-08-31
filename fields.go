package abc

type InformationFieldValue int

const (
	InstructionValue InformationFieldValue = iota
	StringValue
)

type InformationFieldPosition int

const (
	ANYWHERE InformationFieldPosition = iota
	FIRST
	SECOND
	LAST
	NOWHERE
)

type InformationField struct {
	Key                 string
	Abbr                string
	AllowedInFileHeader InformationFieldPosition
	AllowedInTuneHeader InformationFieldPosition
	AllowedInTuneBody   InformationFieldPosition
	Inline              bool
	FieldType           InformationFieldValue
}

var standard21 = map[string]InformationField{
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

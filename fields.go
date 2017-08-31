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
	ValueType           InformationFieldValue
}

// Field represents a InformationField
type Field struct {
	Shadow *InformationField
	Value  string
}

var AreaField = InformationField{"A", "area", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                            //	A:Donegal, A:Bampton (deprecated)
var BookField = InformationField{"B", "book", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                            //	B:O'Neills
var ComposerField = InformationField{"C", "composer", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                    //	C:Robert Jones, C:Trad.
var DiscographyField = InformationField{"D", "discography", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}              //	D:Chieftains IV
var FileField = InformationField{"F", "file", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                            //	F:http://a.b.c/file.abc
var GroupField = InformationField{"G", "group", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                          //	G:flute
var HistoryField = InformationField{"H", "history", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                      //	H:The story behind this tune …
var InstructionField = InformationField{"I", "instruction", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}         // I:papersize A4, I:newpage
var KeyField = InformationField{"K", "key", NOWHERE, LAST, ANYWHERE, true, InstructionValue}                              // K:G, K:Dm, K:AMix
var UnitNoteLengthField = InformationField{"L", "unit note length", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue} // L:1/4, L:1/8
var MeterField = InformationField{"M", "meter", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}                     // M:3/4, M:4/4
var MacroField = InformationField{"m", "macro", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}                     // m: ~G2 = {A},G{F},G
var NotesField = InformationField{"N", "notes", ANYWHERE, ANYWHERE, ANYWHERE, true, StringValue}                          // N:see also O'Neills - 234
var OriginField = InformationField{"O", "origin", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                        // O:UK; Yorkshire; Bradford
var PartsField = InformationField{"P", "parts", NOWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}                      // P:A, P:ABAC, P:(A2B)3
var TempoField = InformationField{"Q", "tempo", NOWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}                      // Q:"allegro" 1/4=120
var RhythmField = InformationField{"R", "rhythm", ANYWHERE, ANYWHERE, ANYWHERE, true, StringValue}                        // R:R, R:reel
var RemarkField = InformationField{"r", "remark", ANYWHERE, ANYWHERE, ANYWHERE, true, StringValue}                        // r:I love abc
var SourceField = InformationField{"S", "source", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}                        // S:collected in Brittany
var SymbolLineField = InformationField{"s", "symbol line", NOWHERE, NOWHERE, ANYWHERE, false, InstructionValue}           // s: !pp! ** !f!
var TuneTitleField = InformationField{"T", "tune title", NOWHERE, SECOND, ANYWHERE, false, StringValue}                   // T:Paddy O'Rafferty
var UserDefinedField = InformationField{"U", "user defined", ANYWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}        // U: T = !trill!
var VoiceField = InformationField{"V", "voice", NOWHERE, ANYWHERE, ANYWHERE, true, InstructionValue}                      // V:4 clef=bass
var WordsField = InformationField{"W", "words", NOWHERE, ANYWHERE, ANYWHERE, false, StringValue}                          // W:lyrics printed after the end of the tune
var WordsLowerField = InformationField{"w", "words", NOWHERE, NOWHERE, ANYWHERE, false, StringValue}                      // w:lyrics printed aligned with the notes of a tune
var ReferenceNumberField = InformationField{"X", "reference number", NOWHERE, FIRST, NOWHERE, false, InstructionValue}    // X:1, X:2
var TranscriptionField = InformationField{"Z", "transcription", ANYWHERE, ANYWHERE, NOWHERE, false, StringValue}          // John, j.s@gmail.com

package abc

type MusicCode string

type TuneHeader struct {
	ReferenceNumber int
	Title           string
	Fields          []InformationField
	Key             string
}

type TuneBody struct {
	Code   MusicCode
	Fields []InformationField
}

type ABCTune struct {
	Body   TuneBody
	Header TuneHeader
	Styles StylesheetDirective
}

type ABCTunebook []ABCTune

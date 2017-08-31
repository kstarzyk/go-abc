package abc

type MusicCode string

type TuneBody struct {
	Code MusicCode
}

type TuneHeader struct {
}

type ABCTune struct {
	Body   TuneBody
	Header TuneHeader
	Styles StylesheetDirective
}

type ABCTunebook []ABCTune

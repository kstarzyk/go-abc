package abc

// An abc fragment is a partial abc tune.
// It may contain a partial tune header with no body or a tune body with optional tune header information fields. (2.3.1)

type ABCFragment struct {
	TuneHeader
	TuneBody
}

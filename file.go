package abc

import (
	"fmt"
	"io/ioutil"
)

type ABCFileStandard float64

type ABCFile struct {
	Version ABCFileStandard
	Tunes   []ABCTune
	Header  FileHeader
	Styles  StylesheetDirective
}

func (f *ABCFile) UnmarshalText(text []byte) error {
	fmt.Println(text)
	return nil
}

func newABCFile() *ABCFile {
	f := &ABCFile{}
	return f
}

func OpenABCFile(path string) (*ABCFile, error) {
	f := newABCFile()

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return f, err
	}

	if err := f.UnmarshalText(content); err != nil {
		return f, err
	}
	return f, nil
}

type FileHeader struct {
	Styles StylesheetDirective
}

type StylesheetDirective string

type TypesetText string

type TextDirective string

package abc

import (
	"reflect"
	"strings"
	"testing"
)

func TestParser_ParseHeader(t *testing.T) {
	var tests = []struct {
		s     string
		stmts []*Field
		err   string
	}{
		// Single field statement
		{
			s: `X:1
T:The Legacy Jig
R:jig
K:G`,
			stmts: []*Field{
				{
					Shadow: Standard21["X"],
					Value:  "1",
				},
				{
					Shadow: Standard21["T"],
					Value:  "The Legacy Jig",
				},
				{
					Shadow: Standard21["R"],
					Value:  "jig",
				},
				{
					Shadow: Standard21["K"],
					Value:  "G",
				},
			},
		},
		{
			s: `X:1
R:jig
T:The Legacy Jig
K:G`,
			stmts: []*Field{
				{
					Shadow: Standard21["X"],
					Value:  "1",
				},
			},
			err: "T must be second declaration",
		},
		{
			s: `X:1
T:The Legacy Jig
R:jig`,
			stmts: []*Field{
				{
					Shadow: Standard21["X"],
					Value:  "1",
				},
				{
					Shadow: Standard21["T"],
					Value:  "The Legacy Jig",
				},
				{
					Shadow: Standard21["R"],
					Value:  "jig",
				},
			},
			err: "K must be last declaration",
		},
	}

	for i, tt := range tests {
		stmts, err := NewParser(strings.NewReader(tt.s)).ParseHeader()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmts, stmts) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmts, stmts)
		}
	}

}

// Ensure the parser can parse strings into Statement ASTs.
func TestParser_ParseField(t *testing.T) {
	var tests = []struct {
		s    string
		stmt *Field
		err  string
	}{
		// Single field statement
		{
			s: `X:1`,
			stmt: &Field{
				Shadow: Standard21["X"],
				Value:  "1",
			},
		},
		{
			s: `M:6/8`,
			stmt: &Field{
				Shadow: Standard21["M"],
				Value:  "6/8",
			},
		},
		{
			s: `T:The Entertainment`,
			stmt: &Field{
				Shadow: Standard21["T"],
				Value:  "The Entertainment",
			},
		},
		{s: `:x`, err: `found ":", expected TEXT`},
	}

	for i, tt := range tests {
		stmt, err := NewParser(strings.NewReader(tt.s)).ParseField()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

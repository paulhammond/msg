package tmpl

import (
	"html/template"
	"testing"
)

func TestXMLDecl(t *testing.T) {
	testFunc(t, "XMLDecl", XMLDecl, testCases{
		{
			name: "no arguments",
			in:   []interface{}{},
			out:  template.HTML(`<?xml version="1.0" encoding="utf-8"?>`),
		},
		{
			name: "more arguments",
			in:   []interface{}{"foo"},
			err:  "too many arguments",
		},
	})
}

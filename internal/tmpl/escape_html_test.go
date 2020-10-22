package tmpl

import (
	"html/template"
	"testing"
)

func TestEscapeHTML(t *testing.T) {
	testFunc(t, "EscapeHTML", EscapeHTML, testCases{
		{
			name: "no arguments",
			in:   []interface{}{},
			err:  "missing str argument",
		},
		{
			name: "simple",
			in:   []interface{}{"Hello"},
			out:  template.HTML("Hello"),
		},
		{
			name: "html",
			in:   []interface{}{`<tag attr="'">&`},
			out:  template.HTML("&lt;tag attr=&#34;&#39;&#34;&gt;&amp;"),
		},
		{
			name: "html as a template.HTML",
			in:   []interface{}{template.HTML(`<tag attr="'">&`)},
			out:  template.HTML("&lt;tag attr=&#34;&#39;&#34;&gt;&amp;"),
		},
		{
			name: "bad html type",
			in:   []interface{}{false},
			err:  "str is not a string",
		},
		{
			name: "more arguments",
			in:   []interface{}{"foo", "bar"},
			err:  "too many arguments",
		},
	})
}

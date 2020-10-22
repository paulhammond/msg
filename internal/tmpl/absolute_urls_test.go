package tmpl

import (
	"testing"
)

func TestAbsoluteURLs(t *testing.T) {
	testFunc(t, "AbsoluteURLs", AbsoluteURLs, testCases{
		{
			name: "simple",
			in:   []interface{}{"https://www.example.com", `<a href="/foo">link</a>`},
			out:  `<a href="https://www.example.com/foo">link</a>`,
		},
		{
			name: "simple",
			in:   []interface{}{"https://www.example.com", `<a href="/foo">link</a>`},
			out:  `<a href="https://www.example.com/foo">link</a>`,
		},
		{
			name: "relative links",
			in:   []interface{}{"https://www.example.com", `<a href="foo">link</a>`},
			out:  `<a href="foo">link</a>`,
		},
		{
			name: "absolute links",
			in:   []interface{}{"https://www.example.com", `<a href="https://www.example.com/foo">link</a>`},
			out:  `<a href="https://www.example.com/foo">link</a>`,
		},
		{
			name: "no arguments",
			in:   []interface{}{},
			err:  "missing html and base argument",
		},
		{
			name: "one argument",
			in:   []interface{}{"hello"},
			err:  "missing html and base argument",
		},
		{
			name: "more arguments",
			in:   []interface{}{"foo", "bar", "baz"},
			err:  "too many arguments",
		},
	})

}

func TestAbsURL(t *testing.T) {
	fragment := `
	this is some text that should survive
	<!-- and this is a comment -->

	<!-- urls that should be adjusted, note the img tag is closed -->
	<p><b><a href="/foo"><img src="/bar"></a></b></p>

	<!-- don't adjust relative urls -->
	<p><b><a href="foo"><img src="bar"></a></b></p>

	<!-- don't adjust protocol relative urls -->
	<p><b><a href="//example.com/foo"><img src="//example.com/bar"></a></b></p>

	<!-- don't adjust absolute urls -->
	<p><b><a href="https://example.com/bar"><img src="https://example.com/bar"></a></b></p>

	<!-- ensure markup that isn't really markup survives even if it is escaped-->
	<pre>href="/foo" &lt;img src="/foo.png"></pre>
	`
	expected := `
	this is some text that should survive
	<!-- and this is a comment -->

	<!-- urls that should be adjusted, note the img tag is closed -->
	<p><b><a href="https://www.example.com/foo"><img src="https://www.example.com/bar"/></a></b></p>

	<!-- don't adjust relative urls -->
	<p><b><a href="foo"><img src="bar"/></a></b></p>

	<!-- don't adjust protocol relative urls -->
	<p><b><a href="//example.com/foo"><img src="//example.com/bar"/></a></b></p>

	<!-- don't adjust absolute urls -->
	<p><b><a href="https://example.com/bar"><img src="https://example.com/bar"/></a></b></p>

	<!-- ensure markup that isn't really markup survives even if it is escaped-->
	<pre>href=&#34;/foo&#34; &lt;img src=&#34;/foo.png&#34;&gt;</pre>
	`

	got, err := absurl(fragment, "https://www.example.com")
	if err != nil {
		t.Errorf("unexpected error %#v", err)
	}
	if got != expected {
		t.Errorf("bad absolute urls\ngot:%#v\nexp:%#v", got, expected)
	}
}

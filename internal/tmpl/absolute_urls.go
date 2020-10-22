package tmpl

import (
	"bytes"
	"errors"
	"reflect"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// AbsoluteURLs converts all path based urls in html to be absolute, as needed
// in Atom and RSS feeds. It expects two arguments, the base url and the HTML.
func AbsoluteURLs(args ...reflect.Value) (interface{}, error) {
	var fragmentv reflect.Value
	var basev reflect.Value

	switch len(args) {
	case 0, 1:
		return nil, errors.New("missing html and base argument")
	case 2:
		basev = args[0]
		fragmentv = args[1]
	default:
		return nil, errors.New("too many arguments")
	}

	fragmentv = indirect(fragmentv)
	basev = indirect(basev)

	// check basev
	if basev.Kind() != reflect.String {
		return nil, errors.New("base is not a string")
	}
	base := basev.String()

	// check fragmentv
	if fragmentv.Kind() != reflect.String {
		return nil, errors.New("fragment is not a string")
	}
	fragment := fragmentv.String()

	return absurl(fragment, base)
}

func absurl(fragment string, base string) (string, error) {
	r := strings.NewReader(fragment)
	nodes, err := html.ParseFragment(r, &html.Node{
		Type:     html.ElementNode,
		Data:     "div",
		DataAtom: atom.Div})
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	for _, n := range nodes {
		absurlWalk(n, base)
		if err := html.Render(&buf, n); err != nil {
			return "", err
		}
	}

	return buf.String(), nil
}

func absurlWalk(n *html.Node, base string) {
	if n.Type == html.ElementNode {
		for i, a := range n.Attr {
			if a.Key == "href" || a.Key == "src" {
				url := a.Val
				if url[0] == '/' && url[1] != '/' {
					n.Attr[i].Val = base + url
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		absurlWalk(c, base)
	}
}

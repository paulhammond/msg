package msg

import (
	"fmt"
	"html/template"
)

type renderContext struct {
	Tree *Tree
	Site tmplv
	Page tmplv
}

// asset takes a path and returns an asset object. Calling this function
// causes a content hash to be added to the path of the requested asset
func (c renderContext) asset(path string) (tmplAsset, error) {
	a, ok := c.Tree.assets[path]
	if !ok {
		return nil, fmt.Errorf("No asset at %q", path)
	}

	err := a.hashify()
	if err != nil {
		return nil, err
	}
	return a.tmplv(), nil

}

// page returns the page being rendered
func (c renderContext) page() tmplv {
	return c.Page
}

// site returns the site object
func (c renderContext) site() tmplv {
	return c.Site
}

func (c renderContext) funcs() template.FuncMap {
	return template.FuncMap{
		"asset": func(s string) (tmplAsset, error) {
			return c.asset(s)
		},
		"page": func() tmplv {
			return c.page()
		},
		"site": func() tmplv {
			return c.site()
		},
	}
}

package msg

import (
	"fmt"
	"html/template"
)

type renderContext struct {
	tree *Tree
}

// asset takes a path and returns an asset object. Calling this function
// causes a content hash to be added to the path of the requested asset
func (c renderContext) asset(path string) (tmplAsset, error) {
	a, ok := c.tree.assets[path]
	if !ok {
		return nil, fmt.Errorf("No asset at %q", path)
	}

	err := a.hashify()
	if err != nil {
		return nil, err
	}
	return a.tmplv(), nil

}

func (c renderContext) funcs() template.FuncMap {
	return template.FuncMap{
		"asset": func(s string) (tmplAsset, error) {
			return c.asset(s)
		},
	}
}

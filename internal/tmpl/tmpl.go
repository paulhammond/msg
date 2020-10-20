package tmpl

import "html/template"

// Map returns the template helpers as a html/template.FuncMap
func Map() template.FuncMap {
	return template.FuncMap{
		"filter":  Filter,
		"sort":    Sort,
		"first":   First,
		"xmldecl": XMLDecl,
	}
}

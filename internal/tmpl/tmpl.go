package tmpl

import "html/template"

// Map returns the template helpers as a html/template.FuncMap
func Map() template.FuncMap {
	return template.FuncMap{
		"absolute_urls": AbsoluteURLs,
		"date_format":   DateFormat,
		"escape_html":   EscapeHTML,
		"filter":        Filter,
		"first":         First,
		"last":          Last,
		"now":           Now,
		"sort":          Sort,
		"xmldecl":       XMLDecl,
	}
}

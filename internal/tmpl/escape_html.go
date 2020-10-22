package tmpl

import (
	"errors"
	"html"
	"html/template"
	"reflect"
)

// EscapeHTML escapes the string using html.EscapeString
func EscapeHTML(args ...reflect.Value) (interface{}, error) {
	var strv reflect.Value

	switch len(args) {
	case 0:
		return nil, errors.New("missing str argument")
	case 1:
		strv = args[0]
	default:
		return nil, errors.New("too many arguments")
	}

	strv = indirect(strv)

	if strv.Kind() != reflect.String {
		return nil, errors.New("str is not a string")
	}
	str := strv.String()

	return template.HTML(html.EscapeString(str)), nil
}

package tmpl

import (
	"errors"
	"html/template"
	"reflect"
)

// XMLDecl outputs the XML declaration <?xml version="1.0" encoding="utf-8"?>.
// We shouldn't really use html/template to generate XML but it works, except
// for the declaration. This function is a workaround for that.
func XMLDecl(args ...reflect.Value) (interface{}, error) {
	if len(args) > 0 {
		return nil, errors.New("too many arguments")
	}
	return template.HTML(`<?xml version="1.0" encoding="utf-8"?>`), nil
}

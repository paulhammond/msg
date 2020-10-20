package tmpl

import (
	"errors"
	"reflect"
	"time"
)

// for testing
var now = time.Now

// Now returns the current time.
func Now(args ...reflect.Value) (interface{}, error) {
	if len(args) > 0 {
		return nil, errors.New("too many arguments")
	}

	return now(), nil
}

package tmpl

import (
	"errors"
	"reflect"
)

// First returns the first count elements of list. It expects either an integer
// count and a slice or array list. The count defaults to 1 if not provided.
func First(args ...reflect.Value) (interface{}, error) {
	var listv reflect.Value
	var countv reflect.Value

	switch len(args) {
	case 0:
		return nil, errors.New("missing list argument")
	case 1:
		countv = reflect.ValueOf(1)
		listv = args[0]
	case 2:
		countv = args[0]
		listv = args[1]
	default:
		return nil, errors.New("too many arguments")
	}

	switch listv.Kind() {
	case reflect.Array, reflect.Slice:
		// all good, continue
	default:
		return nil, errors.New("list is not an array type")
	}

	count, err := valueInt(countv, "count")
	if err != nil {
		return nil, err
	}
	if count > listv.Len() {
		count = listv.Len()
	}
	if count < 0 {
		return nil, errors.New("count must be positive")
	}

	// we tested the list type above
	return listv.Slice(0, count).Interface(), nil
}

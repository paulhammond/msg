package tmpl

import (
	"errors"
	"reflect"
)

// Filter returns any elements of list where the data at key matches value. It
// expects three args, a key, a value and a list. Equality is tested using the
// go == operator.
func Filter(args ...reflect.Value) (interface{}, error) {
	var listv reflect.Value
	var keyv reflect.Value
	var valuev reflect.Value

	switch len(args) {
	case 0, 1, 2:
		return nil, errors.New("missing list, key and value arguments")
	case 3:
		keyv = args[0]
		valuev = args[1]
		listv = args[2]
	default:
		return nil, errors.New("too many arguments")
	}

	listv = indirect(listv)
	keyv = indirect(keyv)
	valuev = indirect(valuev)

	// check list is a list
	list, err := valueList(listv, "list")
	if err != nil {
		return nil, err
	}

	// check sortkey
	if keyv.Kind() != reflect.String {
		return nil, errors.New("key is not a string")
	}
	key := keyv.String()

	var value interface{}

	out := reflect.MakeSlice(reflect.SliceOf(list.Type().Elem()), 0, 0)

	if !valuev.IsValid() || !valuev.CanInterface() {
		return out.Interface(), nil
	}
	value = valuev.Interface()

	for i := 0; i < list.Len(); i++ {
		e := list.Index(i)
		v, err := lookup(e, key)
		if err != nil {
			return nil, err
		}
		if !v.IsValid() {
			continue
		}
		if v.Interface() == value {
			out = reflect.Append(out, e)
		}
	}

	return out.Interface(), nil

}

package tmpl

import (
	"fmt"
	"reflect"
	"strings"
)

// valueInt returns v as an integer, or an error if v is not an integer type.
// This function is intended for calculating array indexes and probably has bugs
// with large unsigned integers or 64 bit integers on a 32 bit platform.
func valueInt(v reflect.Value, name string) (int, error) {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(v.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return int(v.Uint()), nil
	default:
		return 0, fmt.Errorf("%s is not an integer", name)
	}
}

// valueList returns v as a slice or array. If v is a map it returns the values
// of v. If v is a slice or array it returns v. If v is another type it returns
// an error.
func valueList(v reflect.Value, name string) (reflect.Value, error) {
	switch v.Kind() {
	case reflect.Map:
		out := reflect.MakeSlice(reflect.SliceOf(v.Type().Elem()), 0, 0)
		for _, k := range v.MapKeys() {
			out = reflect.Append(out, v.MapIndex(k))
		}
		return out, nil
	case reflect.Array, reflect.Slice:
		return v, nil
	default:
		return reflect.Value{}, fmt.Errorf("%s is not an map or array type", name)
	}
}

// indirect resolves pointers and interfaces to get to the underlying value of v
func indirect(v reflect.Value) reflect.Value {
	// this is heavily influenced by text/template in the go standard library
	for ; v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface; v = v.Elem() {
		if v.IsNil() {
			break
		}
	}
	return v
}

// lookup finds the data at key in v. It supports nested . notations for keys.
func lookup(v reflect.Value, key string) (reflect.Value, error) {
	keys := strings.Split(key, ".")

	for _, k := range keys {
		v = indirect(v)
		switch v.Kind() {
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(k))
		case reflect.Struct:
			v = v.FieldByName(k)
		case reflect.Array, reflect.Slice:
			return reflect.Value{}, fmt.Errorf("lookup not yet unimplmented for %s, we should fix this", v.Kind().String()) // TODO
		default:
			return reflect.ValueOf(nil), nil
		}
		if !v.IsValid() {
			return reflect.ValueOf(nil), nil
		}
	}
	return v, nil
}

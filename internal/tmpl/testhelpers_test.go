package tmpl

import (
	"reflect"
	"testing"
)

type tmplFunc func(args ...reflect.Value) (interface{}, error)

type testCases []struct {
	name    string
	in      []interface{}
	out     interface{}
	err     string
	cleaner func(*testing.T, interface{}) interface{}
}

func testFunc(t *testing.T, fname string, f tmplFunc, tests testCases) {
	t.Helper()
	for _, tt := range tests {

		args := make([]reflect.Value, len(tt.in))
		for i, v := range tt.in {
			args[i] = reflect.ValueOf(v)
		}

		got, err := f(args...)

		if tt.cleaner != nil {
			got = tt.cleaner(t, got)
		}
		if tt.err != "" {
			if err == nil || err.Error() != tt.err {
				t.Errorf("%s(%s) unexpected error:\n got:%+v\n exp:%s", fname, tt.name, err, tt.err)
			}
			if got != nil {
				t.Errorf("%s(%s) unexpected value:\n got:%+v\n exp: nil", fname, tt.name, got)
			}
		} else {
			if err != nil {
				t.Errorf("%s(%s) unexpected error:\n got:%+v\n exp: nil", fname, tt.name, err)
			}
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("%s(%s) unexpected value:\n got:%+v\n exp:%+v", fname, tt.name, got, tt.out)
			}
		}
	}
}

package tmpl

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	a1 := makeStructure("a", 1)
	a2 := makeStructure("a", 2)
	a3 := makeStructure("a", 3)
	b1 := makeStructure("b", 1)
	c1 := makeStructure("c", 1)
	testSlice := []interface{}{a1, a2, a3, b1, c1}
	testMap := map[string]interface{}{"a1": a1, "a2": a2, "a3": a3, "b1": b1, "c1": c1}

	var keyStr = "Key"
	var aStr = "a"

	testFunc(t, "Filter", Filter, testCases{
		{
			name: "slice",
			in:   []interface{}{"Int", 1, testSlice},
			out:  []interface{}{a1, b1, c1},
		},
		{
			name: "map",
			in:   []interface{}{"Int", 1, testMap},
			out:  []interface{}{a1, b1, c1},
			cleaner: func(t *testing.T, i interface{}) interface{} {
				i, err := Sort(reflect.ValueOf("Key"), reflect.ValueOf(i))
				if err != nil {
					t.Errorf("unexpected error %v", err)
				}
				return i
			},
		},
		{
			name: "slice strings",
			in:   []interface{}{"Key", "a", testSlice},
			out:  []interface{}{a1, a2, a3},
		},
		{
			name: "sliceref",
			in:   []interface{}{"Key", "a", &testSlice},
			out:  []interface{}{a1, a2, a3},
		},
		{
			name: "keyref",
			in:   []interface{}{&keyStr, "a", testSlice},
			out:  []interface{}{a1, a2, a3},
		},
		{
			name: "valueref",
			in:   []interface{}{&keyStr, &aStr, testSlice},
			out:  []interface{}{a1, a2, a3},
		},
		{
			name: "missing key",
			in:   []interface{}{"notfound", "a", testSlice},
			out:  []interface{}{},
		},
		{
			name: "sometimes missing key",
			in:   []interface{}{"Strings.a", "a", testSlice},
			out:  []interface{}{a1, a2, a3},
		},
		{
			name: "nil value matches nothing",
			in:   []interface{}{"a", nil, testSlice},
			out:  []interface{}{},
		},
		{
			name: "no args",
			in:   []interface{}{},
			err:  "missing list, key and value arguments",
		},
		{
			name: "one arg",
			in:   []interface{}{testSlice},
			err:  "missing list, key and value arguments",
		},
		{
			name: "two arga",
			in:   []interface{}{"Key", testSlice},
			err:  "missing list, key and value arguments",
		},
		{
			name: "four args",
			in:   []interface{}{"key", "a", "foo", testSlice},
			err:  "too many arguments",
		},
		{
			name: "bad list",
			in:   []interface{}{"key", "desc", 2},
			err:  "list is not an map or array type",
		},
		{
			name: "bad key",
			in:   []interface{}{1, "value", testSlice},
			err:  "key is not a string",
		},
	})
}

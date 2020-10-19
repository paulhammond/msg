package tmpl

import (
	"testing"
)

func TestSort(t *testing.T) {
	a := makeStructure("a", 1)
	b := makeStructure("b", 2)
	c := makeStructure("c", 3)

	testFunc(t, "Sort", Sort, testCases{
		{
			name: "map of maps",
			in: []interface{}{"key", map[string]interface{}{
				"a": a.Strings,
				"b": b.Strings,
				"c": c.Strings,
			}},
			out: []interface{}{a.Strings, b.Strings, c.Strings},
		},
		{
			name: "slice of maps",
			in: []interface{}{"key", []interface{}{
				c.Strings,
				a.Strings,
				b.Strings,
			}},
			out: []interface{}{a.Strings, b.Strings, c.Strings},
		},
		{
			name: "map of structs",
			in: []interface{}{"Key", map[string]interface{}{
				"b": b,
				"a": a,
				"c": c,
			}},
			out: []interface{}{a, b, c},
		},
		{
			name: "list of structs",
			in:   []interface{}{"Key", []interface{}{b, a, c}},
			out:  []interface{}{a, b, c},
		},
		{
			name: "indirect and nested keys",
			in:   []interface{}{"Submapref.sub.Strings.key", []interface{}{b, a, c}},
			out:  []interface{}{a, b, c},
		},
		{
			name: "asc",
			in: []interface{}{"key", "asc", []interface{}{
				c.Strings,
				a.Strings,
				b.Strings,
			}},
			out: []interface{}{a.Strings, b.Strings, c.Strings},
		},
		{
			name: "desc",
			in: []interface{}{"key", "desc", []interface{}{
				c.Strings,
				a.Strings,
				b.Strings,
			}},
			out: []interface{}{c.Strings, b.Strings, a.Strings},
		},
		{
			name: "non-existent key",
			in:   []interface{}{"notfound", []interface{}{c, a, b}},
			out:  []interface{}{c, a, b},
		},
		{
			name: "mix of existant and non-existant keys",
			// Strings.b is only set on b so that should come first
			in:  []interface{}{"Strings.b", []interface{}{c, a, b}},
			out: []interface{}{b, c, a},
		},
		{
			name: "no args",
			in:   []interface{}{},
			err:  "missing list and sort arguments",
		},
		{
			name: "one arg",
			in:   []interface{}{[]interface{}{a}},
			err:  "missing list and sort arguments",
		},
		{
			name: "four args",
			in:   []interface{}{"key", "desc", "foo", []interface{}{a}},
			err:  "too many arguments",
		},
		{
			name: "bad list",
			in:   []interface{}{"key", "desc", 2},
			err:  "list is not an map or array type",
		},
		{
			name: "bad key",
			in:   []interface{}{1, "desc", []interface{}{a}},
			err:  "key is not a string",
		},
		{
			name: "bad dir",
			in:   []interface{}{"key", true, []interface{}{a}},
			err:  "dir is not 'asc' or 'desc'",
		},
		{
			name: "bad dir",
			in:   []interface{}{"key", "backwards", []interface{}{a}},
			err:  "dir is not 'asc' or 'desc'",
		},
	})
}

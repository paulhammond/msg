package tmpl

import (
	"testing"
)

func TestFirst(t *testing.T) {
	list := []int{1, 2, 3}
	testFunc(t, "First", First, testCases{
		{
			name: "one arg",
			in:   []interface{}{list},
			out:  []int{1},
		},
		{
			name: "count",
			in:   []interface{}{2, list},
			out:  []int{1, 2},
		},
		{
			name: "uint count",
			in:   []interface{}{uint(2), list},
			out:  []int{1, 2},
		},
		{
			name: "count > len",
			in:   []interface{}{4, list},
			out:  []int{1, 2, 3},
		},
		{
			name: "zero count",
			in:   []interface{}{0, list},
			out:  []int{},
		},
		{
			name: "empty list with a count",
			in:   []interface{}{1, []int{}},
			out:  []int{},
		},
		{
			name: "empty list with zero count",
			in:   []interface{}{0, []int{}},
			out:  []int{},
		},
		{
			name: "bad list",
			in:   []interface{}{list, 2},
			err:  "list is not an array type",
		},
		{
			name: "bad count",
			in:   []interface{}{"two", list},
			err:  "count is not an integer",
		},
		{
			name: "no args",
			in:   []interface{}{},
			err:  "missing list argument",
		},
		{
			name: "three args",
			in:   []interface{}{2, 3, list},
			err:  "too many arguments",
		},
	})
}

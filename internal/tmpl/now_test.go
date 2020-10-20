package tmpl

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t1 := time.Date(2020, 10, 20, 19, 55, 9, 0, time.UTC)
	now = func() time.Time {
		return t1
	}

	testFunc(t, "Now", Now, testCases{
		{
			name: "no arguments",
			in:   []interface{}{},
			out:  t1,
		},
		{
			name: "more arguments",
			in:   []interface{}{"foo"},
			err:  "too many arguments",
		},
	})
}

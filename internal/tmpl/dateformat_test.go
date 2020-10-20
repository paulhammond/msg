package tmpl

import (
	"testing"
	"time"
)

func TestDateFormat(t *testing.T) {
	testFunc(t, "DateFormat", DateFormat, testCases{
		{
			name: "string arg",
			in:   []interface{}{"02 Jan 2006 15:04 -0700", "2020-10-20T12:55:09-07:00"},
			out:  "20 Oct 2020 12:55 -0700",
		},
		{
			name: "string arg utc",
			in:   []interface{}{"02 Jan 2006 15:04 -0700", "2020-10-20T19:55:09Z"},
			out:  "20 Oct 2020 19:55 +0000",
		},
		{
			name: "time date arg",
			in:   []interface{}{"02 Jan 2006 15:04 -0700", time.Date(2020, 10, 20, 19, 55, 9, 0, time.UTC)},
			out:  "20 Oct 2020 19:55 +0000",
		},
		{
			name: "weird format string",
			in:   []interface{}{"foo", "2020-10-20T19:55:09Z"},
			out:  "foo",
		},
		{
			name: "bad date type",
			in:   []interface{}{"02 Jan 2006 15:04 -0700", 3},
			err:  "date is not a date or string",
		},
		{
			name: "bad date format",
			in:   []interface{}{"02 Jan 2006 15:04 -0700", "foo"},
			err:  `parsing time "foo" as "2006-01-02T15:04:05Z07:00": cannot parse "foo" as "2006"`,
		},
		{
			name: "bad format type",
			in:   []interface{}{false, "2020-10-20T19:55:09Z"},
			err:  "format is not a string",
		},
		{
			name: "no args",
			in:   []interface{}{},
			err:  "missing date and format arguments",
		},
		{
			name: "one arg",
			in:   []interface{}{"2020-10-20T19:55:09Z"},
			err:  "missing date and format arguments",
		},
		{
			name: "three args",
			in:   []interface{}{"02 Jan 2006 15:04 -0700", "2020-10-20T19:55:09Z", "foo"},
			err:  "too many arguments",
		},
	})
}

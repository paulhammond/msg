package tmpl

import (
	"errors"
	"reflect"
	"time"
)

// DateFormat formats the date according to the format. It expects exactly two
// arguments, a format (which follows the patterns used by Go's Time.Format) and
// a date, which can either be a date or a string in RFC3339 format, or a date
// object
func DateFormat(args ...reflect.Value) (interface{}, error) {
	var formatv reflect.Value
	var datev reflect.Value

	switch len(args) {
	case 0, 1:
		return nil, errors.New("missing date and format arguments")
	case 2:
		formatv = args[0]
		datev = args[1]
	default:
		return nil, errors.New("too many arguments")
	}

	formatv = indirect(formatv)
	datev = indirect(datev)

	// check formatv
	if formatv.Kind() != reflect.String {
		return nil, errors.New("format is not a string")
	}
	format := formatv.String()

	// check datev
	var date time.Time
	var err error
	switch datet := datev.Interface().(type) {
	case string:
		date, err = time.Parse(time.RFC3339, datev.String())
	case time.Time:
		date = datet
	default:
		return nil, errors.New("date is not a date or string")
	}
	if err != nil {
		return nil, err
	}

	return date.Format(format), nil

}

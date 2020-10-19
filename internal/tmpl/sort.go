package tmpl

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

// Sort sorts a list using the provided key and direction ("asc" or "desc"). If
// a direction is not provided it defaults to "asc". Args are either "key,
// direction, list" or just "key, direction"
func Sort(args ...reflect.Value) (interface{}, error) {
	var listv reflect.Value
	var keyv reflect.Value
	var dirv reflect.Value

	switch len(args) {
	case 0, 1:
		return nil, errors.New("missing list and sort arguments")
	case 2:
		keyv = args[0]
		dirv = reflect.ValueOf("asc") // TODO: should this be desc?
		listv = args[1]
	case 3:
		keyv = args[0]
		dirv = args[1]
		listv = args[2]
	default:
		return nil, errors.New("too many arguments")
	}

	listv = indirect(listv)
	keyv = indirect(keyv)
	dirv = indirect(dirv)

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

	// check direction
	if dirv.Kind() != reflect.String {
		return nil, errors.New("dir is not 'asc' or 'desc'")
	}
	dir := dirv.String()
	if !(dir == "asc" || dir == "desc") {
		return nil, errors.New("dir is not 'asc' or 'desc'")
	}

	// prep a list to do the sort on
	sortee := []sortElem{}
	for i := 0; i < list.Len(); i++ {
		e := list.Index(i)
		sorter, err := lookup(e, key)
		if err != nil {
			return nil, err
		}
		sortee = append(sortee, sortElem{
			sorter: sorter,
			value:  e,
		})
	}

	err = sortList(sortee, dir == "desc")
	if err != nil {
		return nil, err
	}

	out := reflect.MakeSlice(reflect.SliceOf(list.Type().Elem()), len(sortee), len(sortee))

	for i, v := range sortee {
		out.Index(i).Set(v.value)
	}
	return out.Interface(), nil
}

func sortList(l []sortElem, desc bool) (err error) {
	defer func() {
		if panicErr := recover(); err == nil && panicErr != nil {
			if e2, ok := panicErr.(error); ok {
				err = e2
			}
		}
	}()

	sort.SliceStable(l, func(i, j int) bool {
		if desc {
			i, j = j, i
		}
		r, err := less(l[i].sorter, l[j].sorter)
		if err != nil {
			panic(err)
		}
		return r
	})

	return nil

}

func less(a, b reflect.Value) (bool, error) {
	a = indirect(a)
	b = indirect(b)

	// if a is invalid then it should not sort before b (even if b is invalid too)
	if !a.IsValid() {
		return false, nil
	}
	// if b is invalid it should sort after everything except invalid elements (which is covered above)
	if !b.IsValid() {
		return true, nil
	}

	// TODO: support values that aren't strings
	if a.Kind() == reflect.String && b.Kind() == reflect.String {
		return a.String() < b.String(), nil
	}
	return false, fmt.Errorf("cannot sort %s", a.Kind().String())
}

type sortElem struct {
	sorter reflect.Value
	value  reflect.Value
}

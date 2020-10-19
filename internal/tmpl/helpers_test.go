package tmpl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLookup(t *testing.T) {
	s := makeStructure("a", 1)

	overlyindirected := interface{}(&s)
	tests := []struct {
		obj   interface{}
		key   string
		value interface{}
		err   string
	}{
		{
			obj:   s,
			key:   "Key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Interface",
			value: "a",
		},
		{
			obj:   s,
			key:   "Strings.key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Strings.key2",
			value: "a2",
		},
		{
			obj:   s,
			key:   "Ints.int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Ints.int2",
			value: 3,
		},
		{
			obj:   s,
			key:   "Interfaces.key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Interfaces.int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Sub.Key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Sub.Int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Sub.Strings.key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Sub.Ints.int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Subref.Key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Subref.Int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Subinterface.Key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Subinterface.Int",
			value: 1,
		},
		{
			obj:   s,
			key:   "Submap.sub.Key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Submapref.sub.Key",
			value: "a",
		},
		{
			obj:   s,
			key:   "Submapinterface.sub.Key",
			value: "a",
		},
		{
			obj:   &s,
			key:   "Key",
			value: "a",
		},
		{
			obj:   interface{}(s),
			key:   "Key",
			value: "a",
		},
		{
			obj:   interface{}(&s),
			key:   "Key",
			value: "a",
		},
		{
			obj:   interface{}(&overlyindirected),
			key:   "Key",
			value: "a",
		},
		{
			obj:   s.Strings,
			key:   "key",
			value: "a",
		},
		{
			obj:   &s.Strings,
			key:   "key",
			value: "a",
		},
		{
			obj:   interface{}(s.Strings),
			key:   "key",
			value: "a",
		},
		{
			obj:   s.Ints,
			key:   "int",
			value: 1,
		},
		{
			obj:   s.Interfaces,
			key:   "key",
			value: "a",
		},
		{
			obj:   s,
			key:   "notfound",
			value: nil,
		},
		{
			obj:   s,
			key:   "notfound.key",
			value: nil,
		},
		{
			obj:   s,
			key:   "strings.notfound",
			value: nil,
		},
		{
			obj:   s,
			key:   "sub.notfound",
			value: nil,
		},
		{
			obj:   s,
			key:   "Key.notamap",
			value: nil,
		},
		{
			obj:   s,
			key:   "Int.notamap",
			value: nil,
		},
		{
			obj:   "string",
			key:   "anything",
			value: nil,
		},
		{
			obj:   nil,
			key:   "anything",
			value: nil,
		},
	}

	for _, tt := range tests {
		gotv, err := lookup(reflect.ValueOf(tt.obj), tt.key)
		var got interface{}
		if gotv.IsValid() && gotv.CanInterface() {
			got = gotv.Interface()
		}
		if tt.err != "" {
			if err == nil || err.Error() != tt.err {
				t.Fatalf("lookup(%s) error:\n got:%#v\n exp:%s", tt.key, err, tt.err)
			}
			if got != nil {
				t.Errorf("lookup(%s) value:\n got:%#v\n exp: nil", tt.key, got)
			}
		} else {
			if err != nil {
				t.Fatalf("lookup(%s) error:\n got:%#v\n exp: nil", tt.key, err)
			}
			if !reflect.DeepEqual(got, tt.value) {
				t.Errorf("lookup(%s)  value:\n got:%#v\n exp:%#v", tt.key, got, tt.value)
			}
		}
	}
}

// structure is a big struct that contains many variations of nested maps and
// structs with varying amounts of indirection
type structure struct {
	Key             string
	Int             int
	Interface       interface{}
	Strings         map[string]string
	Ints            map[string]int
	Interfaces      map[string]interface{}
	Sub             substructure
	Subref          *substructure
	Subinterface    interface{}
	Submap          map[string]substructure
	Submapref       map[string]*substructure
	Submapinterface map[string]interface{}
}

// String implements fmt.Stringer so test failures are readable
func (s structure) String() string {
	return fmt.Sprintf("{%s %d}", s.Key, s.Int)
}

// substructure is a component of structure
type substructure struct {
	Key     string
	Int     int
	Strings map[string]string
	Ints    map[string]int
}

// makeStructure returns a structure
func makeStructure(s string, i int) structure {
	sub := substructure{
		Key: s,
		Int: i,

		Strings: map[string]string{
			"key":  s,
			"key2": s + "2",
		},
		Ints: map[string]int{
			"int":  i,
			"int2": i + 2,
		},
	}

	return structure{
		Key:       s,
		Int:       i,
		Interface: s,
		Strings: map[string]string{
			"key":  s,
			"key2": s + "2",
			s:      s,
		},
		Ints: map[string]int{
			"int":  i,
			"int2": i + 2,
		},
		Interfaces: map[string]interface{}{
			"key": s,
			"int": i,
		},
		Sub: substructure{
			Key: s,
			Int: i,

			Strings: map[string]string{
				"key":  s,
				"key2": s + "2",
			},
			Ints: map[string]int{
				"int":  i,
				"int2": i + 2,
			},
		},
		Subref:       &sub,
		Subinterface: interface{}(sub),
		Submap: map[string]substructure{
			"sub": sub,
		},
		Submapref: map[string]*substructure{
			"sub": &sub,
		},
		Submapinterface: map[string]interface{}{
			"sub": sub,
		},
	}
}

package equalser

import (
	"reflect"
	"testing"
)

type Equalser interface {
	Equals(interface{}) bool
}

func AssertEquals(t testing.TB, actual, expected interface{}, msg string) {
	if e, ok := actual.(Equalser); ok {
		if !e.Equals(expected) {
			t.Errorf("not equal: %s\nactual=%+v\nexpected=%+v", msg, actual, expected)
		}
		return
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("not equal: %s\nactual=%+v\nexpected=%+v", msg, actual, expected)
	}
}

package reflect

import (
	"reflect"
	"testing"
)

func TestValueof(t *testing.T) {
	a := 100

	typ := reflect.TypeOf(a)
	if typ.Kind() == reflect.Int {
		n2 := reflect.ValueOf(a).Int()
		t.Error(n2 + int64(a))
	}
}

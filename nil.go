package nil

import (
	"fmt"
	"reflect"
)

func Is(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}

func MustNot(i interface{}) {
	if IsNil(i) {
		panic(fmt.Sprintf("%#v (%T) must not be nil", i, i))
	}
}

// convenience functions, if included with .
func ø(i interface{})          { MustNot(i) }
func IsNil(i interface{}) bool { return Is(i) }

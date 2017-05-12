package less

import (
	"reflect"
	"strings"
)

// Less function
type Func func(a, b interface{}) bool

func IntFunc(asc bool) Func {
	if asc {
		return func(a, b interface{}) bool {
			return reflect.ValueOf(a).Int() < reflect.ValueOf(b).Int()
		}
	}
	return func(a, b interface{}) bool {
		return reflect.ValueOf(a).Int() > reflect.ValueOf(b).Int()
	}
}

func FloatFunc(asc bool) Func {
	if asc {
		return func(a, b interface{}) bool {
			return reflect.ValueOf(a).Float() < reflect.ValueOf(b).Float()
		}
	}
	return func(a, b interface{}) bool {
		return reflect.ValueOf(a).Float() > reflect.ValueOf(b).Float()
	}
}

func StringFunc(asc bool) Func {
	if asc {
		return func(a, b interface{}) bool {
			return strings.Compare(a.(string), b.(string)) < 0
		}
	}
	return func(a, b interface{}) bool {
		return strings.Compare(a.(string), b.(string)) > 0
	}
}

package less

import (
	"github.com/itrabbit/go-stp/conversion/float"
	"github.com/itrabbit/go-stp/conversion/integer"
	s "github.com/itrabbit/go-stp/conversion/string"
	"strings"
)

// Less function
type Func func(a, b interface{}) bool

func IntFunc(asc bool) Func {
	if asc {
		return func(a, b interface{}) bool {
			return integer.FromByDef(a, 64, 0) < integer.FromByDef(b, 64, 0)
		}
	}
	return func(a, b interface{}) bool {
		return integer.FromByDef(a, 64, 0) > integer.FromByDef(b, 64, 0)
	}
}

func FloatFunc(asc bool) Func {
	if asc {
		return func(a, b interface{}) bool {
			return float.FromByDef(a, 64, 0) < float.FromByDef(b, 64, 0)
		}
	}
	return func(a, b interface{}) bool {
		return float.FromByDef(a, 64, 0) > float.FromByDef(b, 64, 0)
	}
}

func StringFunc(asc bool) Func {
	if asc {
		return func(a, b interface{}) bool {
			return strings.Compare(s.FromByDef(a, ""), s.FromByDef(b, "")) < 0
		}
	}
	return func(a, b interface{}) bool {
		return strings.Compare(s.FromByDef(a, ""), s.FromByDef(b, "")) > 0
	}
}

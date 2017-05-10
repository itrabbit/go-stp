package integer

import (
	"reflect"
)

func Type(bitSize int, unsigned bool) reflect.Type {
	switch bitSize {
	case 8:
		if unsigned {
			return reflect.TypeOf(uint8(0))
		} else {
			return reflect.TypeOf(int8(0))
		}
	case 16:
		if unsigned {
			return reflect.TypeOf(uint16(0))
		} else {
			return reflect.TypeOf(int16(0))
		}
	case 32:
		if unsigned {
			return reflect.TypeOf(uint32(0))
		} else {
			return reflect.TypeOf(int32(0))
		}
	case 64:
		if unsigned {
			return reflect.TypeOf(uint64(0))
		} else {
			return reflect.TypeOf(int64(0))
		}
	default:
		if unsigned {
			return reflect.TypeOf(uint(0))
		} else {
			return reflect.TypeOf(int(0))
		}
	}
}

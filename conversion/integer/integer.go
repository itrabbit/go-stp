package integer

import (
	"github.com/itrabbit/go-stp/conversion"
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

// Convert To Int
func To(obj interface{}, bitSize int) (int64, error) {
	t := reflect.TypeOf(obj)
	if t == Type(bitSize, false) {
		return reflect.ValueOf(obj).Int(), nil
	}
	m, err := conversion.GetMethod(t, Type(bitSize, false))
	if err != nil {
		return 0, err
	}
	res, err := m(obj)
	if err != nil {
		return 0, err
	}
	return reflect.ValueOf(res).Int(), nil
}

// Convert To Float By Default Value
func ToDef(obj interface{}, bitSize int, def int64) int64 {
	res, err := To(obj, bitSize)
	if err != nil {
		return def
	}
	return res
}

// Convert To Int Unsigned
func ToUnsigned(obj interface{}, bitSize int) (uint64, error) {
	t := reflect.TypeOf(obj)
	if t == Type(bitSize, true) {
		return reflect.ValueOf(obj).Uint(), nil
	}
	m, err := conversion.GetMethod(t, Type(bitSize, true))
	if err != nil {
		return 0, err
	}
	res, err := m(obj)
	if err != nil {
		return 0, err
	}
	return reflect.ValueOf(res).Uint(), nil
}

// Convert To Float By Default Value
func ToUnsignedDef(obj interface{}, bitSize int, def uint64) uint64 {
	res, err := ToUnsigned(obj, bitSize)
	if err != nil {
		return def
	}
	return res
}

func init() {
	signed := []reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)),
	}
	unsigned := []reflect.Type{
		reflect.TypeOf(uint(0)),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
	}
	sToS := func(obj interface{}) (interface{}, error) {
		return reflect.ValueOf(obj).Int(), nil
	}
	sToU := func(obj interface{}) (interface{}, error) {
		return uint64(reflect.ValueOf(obj).Int()), nil
	}
	uToU := func(obj interface{}) (interface{}, error) {
		return reflect.ValueOf(obj).Uint(), nil
	}
	uToS := func(obj interface{}) (interface{}, error) {
		return int64(reflect.ValueOf(obj).Uint()), nil
	}
	for _, from := range signed {
		for _, to := range signed {
			if from != to {
				conversion.SetMethod(from, to, sToS)
			}
		}
		for _, to := range unsigned {
			conversion.SetMethod(from, to, sToU)
		}
	}
	for _, from := range unsigned {
		for _, to := range unsigned {
			if from != to {
				conversion.SetMethod(from, to, uToU)
			}
		}
		for _, to := range signed {
			conversion.SetMethod(from, to, uToS)
		}
	}
}

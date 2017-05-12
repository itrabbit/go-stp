package integer

import (
	"errors"
	"github.com/itrabbit/go-stp/conversion"
	"reflect"
	"strconv"
	"time"
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
func From(obj interface{}, bitSize int) (int64, error) {
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
func FromByDef(obj interface{}, bitSize int, def int64) int64 {
	res, err := From(obj, bitSize)
	if err != nil {
		return def
	}
	return res
}

// Convert To Int Unsigned
func UnsignedFrom(obj interface{}, bitSize int) (uint64, error) {
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
func UnsignedFromByDef(obj interface{}, bitSize int, def uint64) uint64 {
	res, err := UnsignedFrom(obj, bitSize)
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
	float32ToS := func(obj interface{}) (interface{}, error) {
		return int64(float32(reflect.ValueOf(obj).Float())), nil
	}
	float32ToU := func(obj interface{}) (interface{}, error) {
		return uint64(float32(reflect.ValueOf(obj).Float())), nil
	}
	float64ToS := func(obj interface{}) (interface{}, error) {
		return int64(reflect.ValueOf(obj).Float()), nil
	}
	float64ToU := func(obj interface{}) (interface{}, error) {
		return uint64(reflect.ValueOf(obj).Float()), nil
	}
	boolToS := func(obj interface{}) (interface{}, error) {
		if reflect.ValueOf(obj).Bool() {
			return int64(1), nil
		}
		return int64(0), nil
	}
	boolToU := func(obj interface{}) (interface{}, error) {
		if reflect.ValueOf(obj).Bool() {
			return uint64(1), nil
		}
		return uint64(0), nil
	}
	timeToS := func(obj interface{}) (interface{}, error) {
		if t, ok := obj.(time.Time); ok {
			return t.Unix(), nil
		}
		return nil, errors.New("Conversion object is not time.Time")
	}
	timeToU := func(obj interface{}) (interface{}, error) {
		if t, ok := obj.(time.Time); ok {
			return uint64(t.Unix()), nil
		}
		return nil, errors.New("Conversion object is not time.Time")
	}
	strToS := func(obj interface{}) (interface{}, error) {
		if s, ok := obj.(string); ok {
			return strconv.ParseInt(s, 10, 64)
		}
		return nil, errors.New("Conversion object is not string")

		return strconv.FormatInt(reflect.ValueOf(obj).Int(), 10), nil
	}
	strToU := func(obj interface{}) (interface{}, error) {
		if s, ok := obj.(string); ok {
			return strconv.ParseUint(s, 10, 64)
		}
		return nil, errors.New("Conversion object is not string")
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
		conversion.SetMethod(reflect.TypeOf(string("")), from, strToS)
		conversion.SetMethod(reflect.TypeOf(time.Time{}), from, timeToS)
		conversion.SetMethod(reflect.TypeOf(bool(false)), from, boolToS)
		conversion.SetMethod(reflect.TypeOf(float32(0)), from, float32ToS)
		conversion.SetMethod(reflect.TypeOf(float64(0)), from, float64ToS)
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
		conversion.SetMethod(reflect.TypeOf(string("")), from, strToU)
		conversion.SetMethod(reflect.TypeOf(time.Time{}), from, timeToU)
		conversion.SetMethod(reflect.TypeOf(bool(false)), from, boolToU)
		conversion.SetMethod(reflect.TypeOf(float32(0)), from, float32ToU)
		conversion.SetMethod(reflect.TypeOf(float64(0)), from, float64ToU)
	}
}

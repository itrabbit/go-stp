package float

import (
	"errors"
	"github.com/itrabbit/go-stp/conversion"
	"reflect"
	"strconv"
	"time"
)

// Get Float reflect Type
func Type(bitSize int) reflect.Type {
	switch bitSize {
	case 32:
		return reflect.TypeOf(float32(0))
	case 64:
		return reflect.TypeOf(float64(0))
	default:
		return reflect.TypeOf(float32(0))
	}
}

// Convert To Float
func From(obj interface{}, bitSize int) (float64, error) {
	t := reflect.TypeOf(obj)
	if t == Type(bitSize) {
		return reflect.ValueOf(obj).Float(), nil
	}
	m, err := conversion.GetMethod(t, Type(bitSize))
	if err != nil {
		return 0, err
	}
	res, err := m(obj)
	if err != nil {
		return 0, err
	}
	return reflect.ValueOf(res).Float(), nil
}

// Convert To Float By Default Value
func FromByDef(obj interface{}, bitSize int, def float64) float64 {
	res, err := From(obj, bitSize)
	if err != nil {
		return def
	}
	return res
}

func init() {
	// Float 32 To Float 64
	conversion.SetMethod(Type(32), Type(64), func(obj interface{}) (interface{}, error) {
		if f, ok := obj.(float32); ok {
			return float64(f), nil
		}
		return nil, errors.New("Conversion object is not float32")
	})

	// Float 64 To Float 32
	conversion.SetMethod(Type(64), Type(32), func(obj interface{}) (interface{}, error) {
		if f, ok := obj.(float64); ok {
			return float32(f), nil
		}
		return nil, errors.New("Conversion object is not float64")
	})

	// Signed Integers to Float 32
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)),
	}, Type(32), func(obj interface{}) (interface{}, error) {
		return float32(reflect.ValueOf(obj).Int()), nil
	})

	// Unsigned Integers to Float 32
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(uint(0)),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
	}, Type(32), func(obj interface{}) (interface{}, error) {
		return float32(reflect.ValueOf(obj).Uint()), nil
	})

	// Signed Integers to Float 64
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)),
	}, Type(64), func(obj interface{}) (interface{}, error) {
		return float64(reflect.ValueOf(obj).Int()), nil
	})

	// Unsigned Integers to Float 64
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(uint(0)),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
	}, Type(64), func(obj interface{}) (interface{}, error) {
		return float64(reflect.ValueOf(obj).Uint()), nil
	})

	// String To Float 32
	conversion.SetMethod(reflect.TypeOf(string("")), Type(32), func(obj interface{}) (interface{}, error) {
		if s, ok := obj.(string); ok {
			f, err := strconv.ParseFloat(s, 32)
			if err != nil {
				return nil, err
			}
			return float32(f), nil
		}
		return nil, errors.New("Conversion object is not string")
	})

	// String To Float 64
	conversion.SetMethod(reflect.TypeOf(string("")), Type(64), func(obj interface{}) (interface{}, error) {
		if s, ok := obj.(string); ok {
			return strconv.ParseFloat(s, 64)
		}
		return nil, errors.New("Conversion object is not string")
	})

	// Bool to Float 32
	conversion.SetMethod(reflect.TypeOf(bool(false)), Type(32), func(obj interface{}) (interface{}, error) {
		if b, ok := obj.(bool); ok {
			if b {
				return float32(1.0), nil
			}
			return float32(0.0), nil
		}
		return nil, errors.New("Conversion object is not bool")
	})

	// Bool to Float 64
	conversion.SetMethod(reflect.TypeOf(bool(false)), Type(32), func(obj interface{}) (interface{}, error) {
		if b, ok := obj.(bool); ok {
			if b {
				return float64(1.0), nil
			}
			return float64(0.0), nil
		}
		return nil, errors.New("Conversion object is not bool")
	})

	// time.Time To Float 32
	conversion.SetMethod(reflect.TypeOf(time.Time{}), Type(32), func(obj interface{}) (interface{}, error) {
		if t, ok := obj.(time.Time); ok {
			return float32(t.Unix()), nil
		}
		return nil, errors.New("Conversion object is not time.Time")
	})

	// time.Time To Float 64
	conversion.SetMethod(reflect.TypeOf(time.Time{}), Type(64), func(obj interface{}) (interface{}, error) {
		if t, ok := obj.(time.Time); ok {
			return float64(t.Unix()), nil
		}
		return nil, errors.New("Conversion object is not time.Time")
	})
}

package float

import (
	"errors"
	"github.com/itrabbit/go-stp/conversion"
	"reflect"
	"strconv"
	"time"
)

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

// Convert To Float32
func To32(obj interface{}) (float32, error) {
	t := reflect.TypeOf(obj)
	if t == Type(32) {
		return obj.(float32), nil
	}
	m, err := conversion.GetMethod(t, Type(32))
	if err != nil {
		return 0, err
	}
	res, err := m(obj)
	if err != nil {
		return 0, err
	}
	return res.(float32), nil
}

// Convert To Float32 By Default Value
func ToDef32(obj interface{}, def float32) float32 {
	res, err := To32(obj)
	if err != nil {
		return def
	}
	return res
}

// Convert To Float64
func To64(obj interface{}) (float64, error) {
	t := reflect.TypeOf(obj)
	if t == Type(64) {
		return obj.(float64), nil
	}
	m, err := conversion.GetMethod(t, Type(64))
	if err != nil {
		return 0, err
	}
	res, err := m(obj)
	if err != nil {
		return 0, err
	}
	return res.(float64), nil
}

// Convert To Float64 By Default Value
func ToDef64(obj interface{}, def float64) float64 {
	res, err := To64(obj)
	if err != nil {
		return def
	}
	return res
}

// Convert To Float64
func To(obj interface{}) (float64, error) {
	return To64(obj)
}

// Convert To Float64 By Default Value
func ToDef(obj interface{}, def float64) float64 {
	return ToDef64(obj, def)
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
	// Int 8 To Float 32
	conversion.SetMethod(reflect.TypeOf(int8(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int8); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not int8")
	})
	// Int 8 To Float 64
	conversion.SetMethod(reflect.TypeOf(int8(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int8); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not int8")
	})
	// Int 16 To Float 32
	conversion.SetMethod(reflect.TypeOf(int16(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int16); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not int16")
	})
	// Int 16 To Float 64
	conversion.SetMethod(reflect.TypeOf(int16(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int16); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not int16")
	})
	// Int 32 To Float 32
	conversion.SetMethod(reflect.TypeOf(int32(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int32); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not int32")
	})
	// Int 32 To Float 64
	conversion.SetMethod(reflect.TypeOf(int32(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int32); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not int32")
	})
	// Int To Float 32
	conversion.SetMethod(reflect.TypeOf(int(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not int")
	})
	// Int To Float 64
	conversion.SetMethod(reflect.TypeOf(int(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not int")
	})
	// Int 64 To Float 32
	conversion.SetMethod(reflect.TypeOf(int64(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int64); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not int64")
	})
	// Int 64 To Float 64
	conversion.SetMethod(reflect.TypeOf(int64(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int64); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not int64")
	})
	// UInt 8 To Float 32
	conversion.SetMethod(reflect.TypeOf(uint8(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint8); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not uint8")
	})
	// UInt 8 To Float 64
	conversion.SetMethod(reflect.TypeOf(uint8(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint8); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not uint8")
	})
	// UInt 16 To Float 32
	conversion.SetMethod(reflect.TypeOf(uint16(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint16); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not uint16")
	})
	// UInt 16 To Float 64
	conversion.SetMethod(reflect.TypeOf(uint16(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint16); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not uint16")
	})
	// UInt 32 To Float 32
	conversion.SetMethod(reflect.TypeOf(uint32(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint32); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not uint32")
	})
	// UInt 32 To Float 64
	conversion.SetMethod(reflect.TypeOf(uint32(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint32); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not uint32")
	})
	// UInt To Float 32
	conversion.SetMethod(reflect.TypeOf(uint(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not uint")
	})
	// UInt To Float 64
	conversion.SetMethod(reflect.TypeOf(uint(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not uint")
	})
	// UInt 64 To Float 32
	conversion.SetMethod(reflect.TypeOf(uint64(0)), Type(32), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint64); ok {
			return float32(i), nil
		}
		return nil, errors.New("Conversion object is not uint64")
	})
	// UInt 64 To Float 64
	conversion.SetMethod(reflect.TypeOf(uint64(0)), Type(64), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint64); ok {
			return float64(i), nil
		}
		return nil, errors.New("Conversion object is not uint64")
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

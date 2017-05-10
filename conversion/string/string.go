package string

import (
	"errors"
	"reflect"
	"strconv"

	"time"

	"github.com/itrabbit/go-stp/conversion"
)

// Get String reflect Type
func Type() reflect.Type {
	return reflect.TypeOf(string(""))
}

// Convert To String
func To(obj interface{}) (string, error) {
	t := reflect.TypeOf(obj)
	if t == Type() {
		return obj.(string), nil
	}
	m, err := conversion.GetMethod(t, Type())
	if err != nil {
		return "", err
	}
	res, err := m(obj)
	if err != nil {
		return "", err
	}
	return res.(string), nil
}

// Convert To String By Default Value
func ToDef(obj interface{}, def string) string {
	res, err := To(obj)
	if err != nil {
		return def
	}
	return res
}

func init() {
	// Int 8
	conversion.SetMethod(reflect.TypeOf(int8(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int8); ok {
			return strconv.FormatInt(int64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not int8")
	})
	// Int 16
	conversion.SetMethod(reflect.TypeOf(int16(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int16); ok {
			return strconv.FormatInt(int64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not int16")
	})
	// Int 32
	conversion.SetMethod(reflect.TypeOf(int32(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int32); ok {
			return strconv.FormatInt(int64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not int32")
	})
	// Int
	conversion.SetMethod(reflect.TypeOf(int(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int); ok {
			return strconv.FormatInt(int64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not int")
	})
	// Int 64
	conversion.SetMethod(reflect.TypeOf(int64(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(int64); ok {
			return strconv.FormatInt(i, 10), nil
		}
		return nil, errors.New("Conversion object is not int64")
	})
	// UInt 8
	conversion.SetMethod(reflect.TypeOf(uint8(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint8); ok {
			return strconv.FormatUint(uint64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not uint8")
	})
	// UInt 16
	conversion.SetMethod(reflect.TypeOf(uint16(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint16); ok {
			return strconv.FormatUint(uint64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not uint16")
	})
	// UInt 32
	conversion.SetMethod(reflect.TypeOf(uint32(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint32); ok {
			return strconv.FormatUint(uint64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not uint32")
	})
	// UInt
	conversion.SetMethod(reflect.TypeOf(uint(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint); ok {
			return strconv.FormatUint(uint64(i), 10), nil
		}
		return nil, errors.New("Conversion object is not uint")
	})
	// UInt 64
	conversion.SetMethod(reflect.TypeOf(uint64(0)), Type(), func(obj interface{}) (interface{}, error) {
		if i, ok := obj.(uint64); ok {
			return strconv.FormatUint(i, 10), nil
		}
		return nil, errors.New("Conversion object is not uint64")
	})
	// Float32
	conversion.SetMethod(reflect.TypeOf(float32(0)), Type(), func(obj interface{}) (interface{}, error) {
		if f, ok := obj.(float32); ok {
			return strconv.FormatFloat(float64(f), 'f', -1, 32), nil
		}
		return nil, errors.New("Conversion object is not float32")
	})
	// Float64
	conversion.SetMethod(reflect.TypeOf(float64(0)), Type(), func(obj interface{}) (interface{}, error) {
		if f, ok := obj.(float64); ok {
			return strconv.FormatFloat(f, 'f', -1, 64), nil
		}
		return nil, errors.New("Conversion object is not float64")
	})
	// time.Time
	conversion.SetMethod(reflect.TypeOf(time.Time{}), Type(), func(obj interface{}) (interface{}, error) {
		if t, ok := obj.(time.Time); ok {
			return t.Format(time.RFC3339), nil
		}
		return nil, errors.New("Conversion object is not time.Time")
	})
}

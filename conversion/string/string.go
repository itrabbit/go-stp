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

	// Signed Integers
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(int(0)),
		reflect.TypeOf(int8(0)),
		reflect.TypeOf(int16(0)),
		reflect.TypeOf(int32(0)),
		reflect.TypeOf(int64(0)),
	}, Type(), func(obj interface{}) (interface{}, error) {
		return strconv.FormatInt(reflect.ValueOf(obj).Int(), 10), nil
	})

	// Unsigned Integers
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(uint(0)),
		reflect.TypeOf(uint8(0)),
		reflect.TypeOf(uint16(0)),
		reflect.TypeOf(uint32(0)),
		reflect.TypeOf(uint64(0)),
	}, Type(), func(obj interface{}) (interface{}, error) {
		return strconv.FormatUint(reflect.ValueOf(obj).Uint(), 10), nil
	})

	// Floats
	conversion.SetMultiMethod([]reflect.Type{
		reflect.TypeOf(float32(0)),
		reflect.TypeOf(float64(0)),
	}, Type(), func(obj interface{}) (interface{}, error) {
		return strconv.FormatFloat(reflect.ValueOf(obj).Float(), 'f', -1, 32), nil
	})

	// Bool
	conversion.SetMethod(reflect.TypeOf(bool(false)), Type(), func(obj interface{}) (interface{}, error) {
		if b, ok := obj.(bool); ok {
			if b {
				return "true", nil
			}
			return "false", nil
		}
		return nil, errors.New("Conversion object is not bool")
	})

	// time.Time
	conversion.SetMethod(reflect.TypeOf(time.Time{}), Type(), func(obj interface{}) (interface{}, error) {
		if t, ok := obj.(time.Time); ok {
			return t.Format(time.RFC3339), nil
		}
		return nil, errors.New("Conversion object is not time.Time")
	})
}

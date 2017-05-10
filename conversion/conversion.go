package conversion

import (
	"errors"
	"reflect"
	"sync"
)

type Method func(obj interface{}) (interface{}, error)

var cache struct {
	sync.RWMutex
	m map[reflect.Type]map[reflect.Type]Method
}

// Set Conversion Method
func SetMethod(from reflect.Type, to reflect.Type, m Method) {
	cache.Lock()
	if cache.m == nil {
		cache.m = make(map[reflect.Type]map[reflect.Type]Method)
	}
	if _, ok := cache.m[from]; !ok {
		cache.m[from] = make(map[reflect.Type]Method)
	}
	cache.m[from][to] = m
	cache.Unlock()
}

// Set Conversion Method
func SetMultiMethod(from []reflect.Type, to reflect.Type, m Method) {
	cache.Lock()
	if cache.m == nil {
		cache.m = make(map[reflect.Type]map[reflect.Type]Method)
	}
	for _, f := range from {
		if _, ok := cache.m[f]; !ok {
			cache.m[f] = make(map[reflect.Type]Method)
		}
		cache.m[f][to] = m
	}
	cache.Unlock()
}

// Get Conversion Method
func GetMethod(from reflect.Type, to reflect.Type) (Method, error) {
	var result Method

	cache.RLock()
	if cache.m != nil {
		if c, ok := cache.m[from]; ok && c != nil {
			if m, ok := c[to]; ok {
				result = m
			}
		}
	}
	cache.RUnlock()

	if result != nil {
		return result, nil
	}
	return result, errors.New("Method not exist")
}

// Init Package
func init() {
}

package hashmap

import (
	"errors"
	"github.com/itrabbit/go-stp/util"
	"reflect"
)

// HashMap Interface
type Interface interface {
	Keys() []interface{}
	Set(key interface{}, value interface{}) Interface
	Get(key interface{}) (interface{}, error)
	GetByDef(key interface{}, def interface{}) interface{}
	Remove(key interface{}) Interface
	Size() int64
	Data() map[interface{}]interface{}
	Swap(k0, k1 interface{}) Interface
	Clear() Interface
	Begin() util.MapIterator
	End() util.MapIterator
}

// Private class hashmap
type object struct {
	data map[interface{}]interface{}
}

// Private class hashmap iterator
type iterator struct {
	m     *object
	keys  []interface{}
	index int64
}

func (i *iterator) Next() bool {
	if i.index+1 < int64(len(i.keys)) {
		i.index += 1
		return true
	}
	return false
}

func (i *iterator) Prev() bool {
	if i.index-1 >= 0 && len(i.keys) > 0 {
		i.index -= 1
		return true
	}
	return false
}

func (i *iterator) Key() interface{} {
	if i.index >= 0 && i.index < int64(len(i.keys)) {
		return i.keys[i.index]
	}
	return nil
}

func (i *iterator) Data() interface{} {
	if d, err := i.m.Get(i.Key()); err == nil {
		return d
	}
	return nil
}

// Get keys
func (m *object) Keys() []interface{} {
	size := m.Size()
	res := make([]interface{}, size, size)
	index := 0
	for k := range m.data {
		res[index] = k
		index++
	}
	return res
}

// Set value by key
func (m *object) Set(key interface{}, value interface{}) Interface {
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	m.data[key] = value
	return m
}

// Get item by key
func (m *object) Get(key interface{}) (interface{}, error) {
	if item, ok := m.data[key]; ok {
		return item, nil
	}
	return nil, errors.New("Item not fount")
}

// Get item by key
func (m *object) GetByDef(key interface{}, def interface{}) interface{} {
	if item, ok := m.data[key]; ok {
		return item
	}
	return def
}

// Remove item by key
func (m *object) Remove(key interface{}) Interface {
	if _, ok := m.data[key]; ok {
		delete(m.data, key)
	}
	return m
}

// Get MapHash size
func (m *object) Size() int64 {
	return int64(len(m.data))
}

// Get Raw Data
func (m *object) Data() map[interface{}]interface{} {
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	return m.data
}

// Swap items by they keys
func (m *object) Swap(k0, k1 interface{}) Interface {
	m.data[k0], m.data[k1] = m.data[k1], m.data[k0]
	return m
}

// Clear HashMap
func (m *object) Clear() Interface {
	m.data = nil
	m.data = make(map[interface{}]interface{})
	return m
}

// Get begin iterator
func (m *object) Begin() util.MapIterator {
	return &iterator{
		m:     m,
		keys:  m.Keys(),
		index: 0,
	}
}

// Get end iterator
func (m *object) End() util.MapIterator {
	keys := m.Keys()
	return &iterator{
		m:     m,
		keys:  keys,
		index: int64(len(keys) - 1),
	}
}

// Create new hashmap
func New(maps ...interface{}) Interface {
	m := new(object)
	m.data = make(map[interface{}]interface{})
	for _, item := range maps {
		if v := reflect.Indirect(reflect.ValueOf(item)); v.IsValid() && !v.IsNil() {
			if v.Kind() == reflect.Map {
				for _, k := range v.MapKeys() {
					m.Set(k.Interface(), v.MapIndex(k).Interface())
				}
			}
		}
	}
	return m
}

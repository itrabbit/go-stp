package array

import (
	"errors"
)

type Object struct {
	data []interface{}
}

// Set value
func (a *Object) Set(pos int64, obj interface{}) *Object {
	if pos >= 0 && pos < a.Length() {
		a.data[pos] = obj
	}
	return a
}

// Get item from Array by position
func (a *Object) At(pos int64) (interface{}, error) {
	if pos >= 0 && pos < a.Length() {
		return a.data[pos], nil
	}
	return nil, errors.New("Outside the array")
}

// Get first item from Array
func (a *Object) First() (interface{}, error) {
	if a.Length() > 0 {
		return a.data[0], nil
	}
	return nil, errors.New("Array is empty")
}

// Get last item from Array
func (a *Object) Last() (interface{}, error) {
	if l := a.Length(); l-1 >= 0 {
		return a.data[l-1], nil
	}
	return nil, errors.New("Array is empty")
}

// Add object to Array
func (a *Object) Push(obj interface{}) *Object {
	a.data = append(a.data, obj)
	return a
}

// Pop item from Array
func (a *Object) Pop() (interface{}, error) {
	if l := a.Length(); l-1 >= 0 {
		var result interface{}
		result, a.data = a.data[l-1], a.data[:l-1]
		return result, nil
	}
	return nil, errors.New("Array is empty")
}

// Shift item from Array
func (a *Object) Shift() (interface{}, error) {
	if a.Length() > 0 {
		var result interface{}
		result, a.data = a.data[0], a.data[1:]
		return result, nil
	}
	return nil, errors.New("Array is empty")
}

// Insert object to Array by position
func (a *Object) Insert(pos int64, obj interface{}) *Object {
	if pos >= 0 && pos < a.Length() {
		a.data = append(a.data, 0)
		copy(a.data[pos+1:], a.data[pos:])
		a.data[pos] = obj
	} else {
		a.data = append(a.data, obj)
	}
	return a
}

// Remove item by position
func (a *Object) Remove(pos int64) *Object {
	if pos >= 0 && pos < a.Length() {
		copy(a.data[pos:], a.data[pos+1:])
		a.data[len(a.data)-1] = nil
		a.data = a.data[:len(a.data)-1]
	}
	return a
}

// Resize array
func (a *Object) Resize(size int64) *Object {
	if a.Length() > size {
		a.data = a.data[0:size]
	} else if a.Length() < size {
		a.data = append(a.data, make([]interface{}, size-a.Length())...)
	}
	return a
}

// Get Array length
func (a *Object) Length() int64 {
	return int64(len(a.data))
}

// Swap items by they positions
func (a *Object) Swap(i, j int64) *Object {
	a.data[i], a.data[j] = a.data[j], a.data[i]
	return a
}

// Get Raw Data
func (a *Object) Data() []interface{} {
	return a.data
}

// Clear array
func (a *Object) Clear() *Object {
	a.data = a.data[len(a.data):]
	a.data = nil
	return a
}

// Less function
type LessFunc func(a, b interface{}) bool

// Sort Array
func (a *Object) Sort(less LessFunc) *Object {
	n := a.Length()
	maxDepth := int64(0)
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	return a.quickSort(less, 0, n, maxDepth)
}

// Create new array
func New(items ...interface{}) *Object {
	a := new(Object)
	a.data = make([]interface{}, len(items))
	for i, item := range items {
		a.data[i] = item
	}
	return a
}

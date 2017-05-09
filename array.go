package go_stp

import (
	"errors"
)

type Array struct {
	data []interface{}
}

// Get item from Array by position
func (a *Array) At(pos int64) (interface{}, error) {
	if pos >= 0 && pos < a.Length() {
		return a.data[pos], nil
	}
	return nil, errors.New("Outside the array")
}

// Get first item from Array
func (a *Array) First() (interface{}, error) {
	if a.Length() > 0 {
		return a.data[0], nil
	}
	return nil, errors.New("Array is empty")
}

// Get last item from Array
func (a *Array) Last() (interface{}, error) {
	if l := a.Length(); l-1 >= 0 {
		return a.data[l-1], nil
	}
	return nil, errors.New("Array is empty")
}

// Add object to Array
func (a *Array) Push(obj interface{}) *Array {
	a.data = append(a.data, obj)
	return a
}

// Pop item from Array
func (a *Array) Pop() (interface{}, error) {
	if l := a.Length(); l-1 >= 0 {
		var result interface{}
		result, a.data = a.data[l-1], a.data[:l-1]
		return result, nil
	}
	return nil, errors.New("Array is empty")
}

// Shift item from Array
func (a *Array) Shift() (interface{}, error) {
	if a.Length() > 0 {
		var result interface{}
		result, a.data = a.data[0], a.data[1:]
		return result, nil
	}
	return nil, errors.New("Array is empty")
}

// Insert object to Array by position
func (a *Array) Insert(pos int64, obj interface{}) *Array {
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
func (a *Array) Remove(pos int64) *Array {
	if pos >= 0 && pos < a.Length() {
		copy(a.data[pos:], a.data[pos+1:])
		a.data[len(a.data)-1] = nil
		a.data = a.data[:len(a.data)-1]
	}
	return a
}

// Get Array length
func (a *Array) Length() int64 {
	return int64(len(a.data))
}

// Swap items by they positions
func (a *Array) Swap(i, j int64) *Array {
	a.data[i], a.data[j] = a.data[j], a.data[i]
	return a
}

// Get Raw Data
func (a *Array) Data() []interface{} {
	return a.data
}

// Sort Array
func (a *Array) Sort(cmp CompareFunc) *Array {
	n := a.Length()
	maxDepth := int64(0)
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	return a.quickSort(cmp, 0, n, maxDepth)
}

// Create new array
func NewArray(items ...interface{}) *Array {
	arr := new(Array)
	arr.data = make([]interface{}, len(items))
	for i, item := range items {
		arr.data[i] = item
	}
	return arr
}

package array

import (
	"errors"
	"github.com/itrabbit/go-stp/util"
	"github.com/itrabbit/go-stp/util/less"
)

// Array Interface
type Interface interface {
	Set(pos int64, item interface{}) Interface
	Get(pos int64) (interface{}, error)
	GetByDef(pos int64, def interface{}) interface{}
	First() (interface{}, error)
	Last() (interface{}, error)
	Push(items ...interface{}) Interface
	Pop() (interface{}, error)
	Shift() (interface{}, error)
	Insert(pos int64, item interface{}) Interface
	Remove(pos int64) Interface
	Resize(size int64) Interface
	Length() int64
	Swap(i, j int64) Interface
	ReOrder() Interface
	Data() []interface{}
	Begin() util.Iterator
	End() util.Iterator
	Clear() Interface
	Sort(less less.Func) Interface
}

// Private class array
type object struct {
	data []interface{}
}

// Private class array iterator
type iterator struct {
	a     *object
	index int64
}

func (i *iterator) Next() bool {
	if i.index+1 < i.a.Length() {
		i.index += 1
		return true
	}
	return false
}

func (i *iterator) Prev() bool {
	if i.index-1 >= 0 && i.a.Length() > 0 {
		i.index -= 1
		return true
	}
	return false
}

func (i *iterator) Index() int64 {
	return i.index
}

func (i *iterator) Data() interface{} {
	d, err := i.a.Get(i.index)
	if err != nil {
		return nil
	}
	return d
}

// Set value
func (a *object) Set(pos int64, item interface{}) Interface {
	if pos >= 0 && pos < a.Length() {
		a.data[pos] = item
	}
	return a
}

// Get item from Array by position
func (a *object) Get(pos int64) (interface{}, error) {
	if pos >= 0 && pos < a.Length() {
		return a.data[pos], nil
	}
	return nil, errors.New("Outside the array")
}

// Get item from Array by position
func (a *object) GetByDef(pos int64, def interface{}) interface{} {
	if pos >= 0 && pos < a.Length() {
		return a.data[pos]
	}
	return def
}

// Get first item from Array
func (a *object) First() (interface{}, error) {
	if a.Length() > 0 {
		return a.data[0], nil
	}
	return nil, errors.New("Array is empty")
}

// Get last item from Array
func (a *object) Last() (interface{}, error) {
	if l := a.Length(); l-1 >= 0 {
		return a.data[l-1], nil
	}
	return nil, errors.New("Array is empty")
}

// Add objects to Array
func (a *object) Push(items ...interface{}) Interface {
	if a.data == nil {
		a.data = make([]interface{}, 0)
	}
	a.data = append(a.data, items...)
	return a
}

// Pop item from Array
func (a *object) Pop() (interface{}, error) {
	if l := a.Length(); l-1 >= 0 {
		var result interface{}
		result, a.data = a.data[l-1], a.data[:l-1]
		return result, nil
	}
	return nil, errors.New("Array is empty")
}

// Shift item from Array
func (a *object) Shift() (interface{}, error) {
	if a.Length() > 0 {
		var result interface{}
		result, a.data = a.data[0], a.data[1:]
		return result, nil
	}
	return nil, errors.New("Array is empty")
}

// Insert object to Array by position
func (a *object) Insert(pos int64, item interface{}) Interface {
	if a.data == nil {
		a.data = make([]interface{}, 0)
	}
	if pos >= 0 && pos < a.Length() {
		a.data = append(a.data, 0)
		copy(a.data[pos+1:], a.data[pos:])
		a.data[pos] = item
	} else {
		a.data = append(a.data, item)
	}
	return a
}

// Remove item by position
func (a *object) Remove(pos int64) Interface {
	if pos >= 0 && pos < a.Length() {
		copy(a.data[pos:], a.data[pos+1:])
		a.data[len(a.data)-1] = nil
		a.data = a.data[:len(a.data)-1]
	}
	return a
}

// Resize array
func (a *object) Resize(size int64) Interface {
	if a.data == nil {
		a.data = make([]interface{}, 0)
	}
	if a.Length() > size {
		a.data = a.data[0:size]
	} else if a.Length() < size {
		a.data = append(a.data, make([]interface{}, size-a.Length())...)
	}
	return a
}

// Get Array length
func (a *object) Length() int64 {
	return int64(len(a.data))
}

// Swap items by they positions
func (a *object) Swap(i, j int64) Interface {
	a.data[i], a.data[j] = a.data[j], a.data[i]
	return a
}

// ReOrder Array
func (a *object) ReOrder() Interface {
	if a.data == nil {
		a.data = make([]interface{}, 0)
	}
	firstIndex, lastIndex, steps := int64(0), a.Length()-1, a.Length()/2
	if lastIndex > firstIndex {
		for i := int64(0); i < steps; i++ {
			a.Swap(firstIndex+i, lastIndex-i)
		}
	}
	return a
}

// Get Raw Data
func (a *object) Data() []interface{} {
	if a.data == nil {
		a.data = make([]interface{}, 0)
	}
	return a.data
}

// Clear array
func (a *object) Clear() Interface {
	a.data = a.data[len(a.data):]
	a.data = nil
	return a
}

func (a *object) Begin() util.Iterator {
	return &iterator{
		a:     a,
		index: 0,
	}
}

func (a *object) End() util.Iterator {
	index := a.Length() - 1
	if index < 0 {
		index = 0
	}
	return &iterator{
		a:     a,
		index: index,
	}
}

// Sort Array
func (a *object) Sort(less less.Func) Interface {
	if a.data == nil {
		a.data = make([]interface{}, 0)
	}
	n := a.Length()
	maxDepth := int64(0)
	for i := n; i > 0; i >>= 1 {
		maxDepth++
	}
	maxDepth *= 2
	return a.quickSort(less, 0, n, maxDepth)
}

// Create new array
func New(items ...interface{}) Interface {
	a := new(object)
	a.data = make([]interface{}, 0, len(items)+1)
	a.data = append(a.data, items...)
	return a
}

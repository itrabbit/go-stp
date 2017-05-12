package queue

import (
	"github.com/itrabbit/go-stp/array"
)

type Object struct {
	array.Interface
	size int64
}

// Create Queue
func New(size int64, items ...interface{}) *Object {
	q := new(Object)
	q.size = size
	if int64(len(items)) > q.size {
		items = items[0:q.size]
	}
	q.Interface = array.New(items...)
	return q
}

// Resize array
func (q *Object) Resize(size int64) *Object {
	q.size = size
	q.Interface.Resize(q.size)
	return q
}

// Add object to Array
func (q *Object) Insert(pos int64, obj interface{}) *Object {
	if q.size > 0 {
		q.Interface.Insert(pos, obj)
	}
	if q.Length() >= q.size {
		q.Interface.Resize(q.size)
	}
	return q
}

// Add object to Array
func (q *Object) Push(obj interface{}) *Object {
	if q.size > 0 {
		q.Interface.Insert(0, obj)
	}
	if q.Length() >= q.size {
		q.Interface.Resize(q.size)
	}
	return q
}

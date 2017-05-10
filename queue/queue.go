package queue

import (
	"github.com/itrabbit/go-stp/array"
)

type Object struct {
	array.Object
	size int64
}

// Create Queue
func New(size int64, items ...interface{}) *Object {
	q := new(Object)
	q.size = size
	if int64(len(items)) > q.size {
		items = items[0:q.size]
	}
	q.Object = *array.New(items...)
	return q
}

// Resize array
func (q *Object) Resize(size int64) *Object {
	q.size = size
	q.Object.Resize(q.size)
	return q
}

// Add object to Array
func (q *Object) Insert(pos int64, obj interface{}) *Object {
	if q.size > 0 {
		q.Object.Insert(pos, obj)
	}
	if q.Length() >= q.size {
		q.Object.Resize(q.size)
	}
	return q
}

// Add object to Array
func (q *Object) Push(obj interface{}) *Object {
	if q.size > 0 {
		q.Object.Insert(0, obj)
	}
	if q.Length() >= q.size {
		q.Object.Resize(q.size)
	}
	return q
}

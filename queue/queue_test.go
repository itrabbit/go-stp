package queue

import (
	"fmt"
	"testing"
)

func TestObject_Push(t *testing.T) {
	q := New(3, 1, 2, 3, 4)
	if r := fmt.Sprint(q.Push(0).Data()); r != "[0 1 2]" {
		t.Error("Invalid push operation", r, "!=", "[0 1 2]")
	}
}

func TestObject_Insert(t *testing.T) {
	q := New(3, 1, 2, 3)
	if r := fmt.Sprint(q.Insert(1, 2).Data()); r != "[1 2 2]" {
		t.Error("Invalid insert operation", r, "!=", "[1 2 2]")
	}
}

package array

import (
	"fmt"
	"testing"
)

func TestArray_Len(t *testing.T) {
	arr := New(1, 2, 3)
	if arr.Length() != 3 {
		t.Error("Invalid array length")
	}
}

func TestArray_At(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.At(2)
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 3 {
		t.Error("Invalid get item value by position")
	}
}

func TestArray_First(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.First()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 1 {
		t.Error("Invalid first item value")
	}
}

func TestArray_Last(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.Last()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 5 {
		t.Error("Invalid last item value")
	}
}

func TestArray_Push(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.Push(6).Last()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 6 {
		t.Error("Invalid pushed item value")
	}
	if arr.Length() != 6 {
		t.Error("Invalid array length")
	}
}

func TestArray_Pop(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.Pop()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 5 {
		t.Error("Invalid pop item")
	}
	item, err = arr.Last()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 4 {
		t.Error("Invalid last item value")
	}
	if arr.Length() != 4 {
		t.Error("Invalid array length")
	}
}

func TestArray_Shift(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.Shift()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 1 {
		t.Error("Invalid shift item")
	}
	item, err = arr.First()
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 2 {
		t.Error("Invalid first item value")
	}
	if arr.Length() != 4 {
		t.Error("Invalid array length")
	}
}

func TestArray_Insert(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	item, err := arr.Insert(2, 10).At(2)
	if err != nil {
		t.Error(err)
	}
	if item.(int) != 10 {
		t.Error("Invalid insert item")
	}
	if arr.Length() != 6 {
		t.Error("Invalid array length")
	}
}

func TestArray_Remove(t *testing.T) {
	arr := New(1, 2, 3, 4, 5)
	if arr.Remove(2).Length() != 4 {
		t.Error("Invalid array length")
	}
}

func TestArray_Sort(t *testing.T) {
	arr := New(5, 2, 3, 4, 1)
	if fmt.Sprint(arr.Sort(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	}).Data()) != "[5 4 3 2 1]" {
		t.Error("Invalid sort")
	}
}

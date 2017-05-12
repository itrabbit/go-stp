package array

import (
	"fmt"
	"github.com/itrabbit/go-stp/util"
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
	item, err := arr.Get(2)
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
	item, err := arr.Insert(2, 10).Get(2)
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

func TestArray_ForEachNext(t *testing.T) {
	result := ""
	arr := New(5, 2, 3, 4, 1)
	util.Each(arr.Begin(), func(index int64, data interface{}) (stop bool) {
		result += fmt.Sprint(index, data, ";")
		return
	})
	if result != "0 5;1 2;2 3;3 4;4 1;" {
		t.Error("Invalid iterator each next")
	}
}

func TestArray_ForEachBack(t *testing.T) {
	result := ""
	arr := New(5, 2, 3, 4, 1)
	util.EachBack(arr.End(), func(index int64, data interface{}) (stop bool) {
		result += fmt.Sprint(index, data, ";")
		return
	})
	if result != "4 1;3 4;2 3;1 2;0 5;" {
		t.Error("Invalid iterator each back")
	}
}

func TestArray_ReOrder(t *testing.T) {
	arr := New(1, 2, 3, 4, 5, 6, 7, 8, 9)
	result := fmt.Sprint(arr.ReOrder().Data())
	if result != "[9 8 7 6 5 4 3 2 1]" {
		t.Error("Invalid reorder operation")
	}
}

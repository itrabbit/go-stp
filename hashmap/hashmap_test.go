package hashmap

import (
	"fmt"
	"github.com/itrabbit/go-stp/util"
	"github.com/itrabbit/go-stp/util/less"
	"testing"
)

func TestHashMap_Size(t *testing.T) {
	m := New(map[int]int{
		1: 2,
		2: 2,
		3: 2,
	})
	if m.Size() != 3 {
		t.Error("Invalid map size")
	}
}

func TestHashMap_Each(t *testing.T) {
	m := New(map[int]int{
		1: 2,
		2: 2,
		3: 2,
	})
	result := ""
	util.MapEach(m.Begin().SortKeys(less.IntFunc(true)), func(key interface{}, data interface{}) (stop bool) {
		result += fmt.Sprint(key, ":", data, ";")
		return
	})
	if result != "1:2;2:2;3:2;" {
		t.Error("Invalid iterator each next")
	}
}

func TestHashMap_EachBack(t *testing.T) {
	m := New(map[int]int{
		1: 2,
		2: 2,
		3: 2,
	})
	result := ""
	util.MapEachBack(m.End().SortKeys(less.IntFunc(true)), func(key interface{}, data interface{}) (stop bool) {
		result += fmt.Sprint(key, ":", data, ";")
		return
	})
	if result != "3:2;2:2;1:2;" {
		t.Error("Invalid iterator each back")
	}
}

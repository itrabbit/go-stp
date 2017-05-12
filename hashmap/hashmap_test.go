package hashmap

import (
	"github.com/itrabbit/go-stp/util"
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
	size := int64(0)
	util.MapEach(m.Begin(), func(key interface{}, data interface{}) (stop bool) {
		size += 1
		return
	})
	if size != m.Size() {
		t.Error("Invalid iterator each next")
	}
}

func TestHashMap_EachBack(t *testing.T) {
	m := New(map[int]int{
		1: 2,
		2: 2,
		3: 2,
	})
	size := int64(0)
	util.MapEachBack(m.End(), func(key interface{}, data interface{}) (stop bool) {
		size += 1
		return
	})
	if size != m.Size() {
		t.Error("Invalid iterator each back")
	}
}

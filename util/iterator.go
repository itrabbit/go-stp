package util

type Iterator interface {
	Next() bool
	Prev() bool
	Index() int64
	Data() interface{}
}

type MapIterator interface {
	Next() bool
	Prev() bool
	Key() interface{}
	Data() interface{}
}

// Перебор вперед
func Each(it Iterator, callback func(index int64, data interface{}) (stop bool)) {
	if it != nil {
		for true {
			if callback(it.Index(), it.Data()) {
				break
			}
			if !it.Next() {
				break
			}
		}
	}
}

// Перебор вперед
func MapEach(it MapIterator, callback func(key interface{}, data interface{}) (stop bool)) {
	if it != nil {
		for true {
			if callback(it.Key(), it.Data()) {
				break
			}
			if !it.Next() {
				break
			}
		}
	}
}

// Перебор назад
func EachBack(it Iterator, callback func(index int64, data interface{}) (stop bool)) {
	if it != nil {
		for true {
			if callback(it.Index(), it.Data()) {
				break
			}
			if !it.Prev() {
				break
			}
		}
	}
}

// Перебор назад
func MapEachBack(it MapIterator, callback func(key interface{}, data interface{}) (stop bool)) {
	if it != nil {
		for true {
			if callback(it.Key(), it.Data()) {
				break
			}
			if !it.Prev() {
				break
			}
		}
	}
}

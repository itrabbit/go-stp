package array

import (
	"github.com/itrabbit/go-stp/util/less"
)

// Copy and adaptive from "sort" package

func (a *object) quickSort(less less.Func, _a, b, maxDepth int64) *object {
	for b-_a > 12 {
		if maxDepth == 0 {
			return a.heapSort(less, _a, b)
		}
		maxDepth--
		mlo, mhi := a.doPivot(less, _a, b)
		if mlo-_a < b-mhi {
			a.quickSort(less, _a, mlo, maxDepth)
			_a = mhi
		} else {
			a.quickSort(less, mhi, b, maxDepth)
			b = mlo
		}
	}
	if b-_a > 1 {
		for i := _a + 6; i < b; i++ {
			if less(a.data[i], a.data[i-6]) {
				a.Swap(i, i-6)
			}
		}
		a.insertionSort(less, _a, b)
	}
	return a
}

func (a *object) heapSort(less less.Func, _a, b int64) *object {
	first, lo, hi := _a, int64(0), b-_a
	for i := (hi - 1) / 2; i >= 0; i-- {
		a.siftDown(less, i, hi, first)
	}
	for i := hi - 1; i >= 0; i-- {
		a.Swap(first, first+i)
		a.siftDown(less, lo, i, first)
	}
	return a
}

func (a *object) insertionSort(less less.Func, _a, b int64) *object {
	for i := _a + 1; i < b; i++ {
		for j := i; j > _a && less(a.data[j], a.data[j-1]); j-- {
			a.Swap(j, j-1)
		}
	}
	return a
}

func (a *object) siftDown(less less.Func, lo, hi, first int64) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(a.data[first+child], a.data[first+child+1]) {
			child++
		}
		if !less(a.data[first+root], a.data[first+child]) {
			return
		}
		a.Swap(first+root, first+child)
		root = child
	}
}

func (a *object) medianOfThree(less less.Func, m1, m0, m2 int64) {
	if less(a.data[m1], a.data[m0]) {
		a.Swap(m1, m0)
	}
	if less(a.data[m2], a.data[m1]) {
		a.Swap(m2, m1)
		if less(a.data[m1], a.data[m0]) {
			a.Swap(m1, m0)
		}
	}
}

func (a *object) doPivot(less less.Func, lo, hi int64) (midlo, midhi int64) {
	m := lo + (hi-lo)/2
	if hi-lo > 40 {
		s := (hi - lo) / 8
		a.medianOfThree(less, lo, lo+s, lo+2*s)
		a.medianOfThree(less, m, m-s, m+s)
		a.medianOfThree(less, hi-1, hi-1-s, hi-1-2*s)
	}
	a.medianOfThree(less, lo, m, hi-1)

	pivot := lo
	_a, c := lo+1, hi-1

	for ; _a < c && less(a.data[_a], a.data[pivot]); _a++ {
	}
	b := _a
	for {
		for ; b < c && !less(a.data[pivot], a.data[b]); b++ {
		}
		for ; b < c && less(a.data[pivot], a.data[c-1]); c-- {
		}
		if b >= c {
			break
		}
		a.Swap(b, c-1)
		b++
		c--
	}
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		dups := 0
		if !less(a.data[pivot], a.data[hi-1]) {
			a.Swap(c, hi-1)
			c++
			dups++
		}
		if !less(a.data[b-1], a.data[pivot]) {
			b--
			dups++
		}
		if !less(a.data[m], a.data[pivot]) {
			a.Swap(m, b-1)
			b--
			dups++
		}
		protect = dups > 1
	}
	if protect {
		for {
			for ; _a < b && !less(a.data[b-1], a.data[pivot]); b-- {
			}
			for ; _a < b && less(a.data[_a], a.data[pivot]); _a++ {
			}
			if _a >= b {
				break
			}
			a.Swap(_a, b-1)
			_a++
			b--
		}
	}
	a.Swap(pivot, b-1)
	return b - 1, c
}

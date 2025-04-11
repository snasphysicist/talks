package fallback

import (
	"iter"
	"math/rand"
)

func setsMain() {
	smallSlice := []int{1, 2, 3, 4, 5, 4, 6, 2, 3, 10, 4, 6, 7, 8, 8, 9}
	println("SMALL SLICE")
	for _, e := range NewSet(smallSlice).Elements() {
		println("Eager", e)
	}
	for e := range LazySetIterator(smallSlice) {
		println("Lazy", e)
	}
	println("MAKING BIG SLICE")
	bigSlice := lotsOfItems()
	println("MADE BIG SLICE")
	println("START RANGE-ING EAGER")
	for _, e := range NewSet(bigSlice).Elements() {
		println("Eager", e)
	}
	println("START RANGE-ING LAZY")
	for e := range LazySetIterator(bigSlice) {
		println("Lazy", e)
	}
}

type Set[T comparable] struct {
	es []T
}

func NewSet[T comparable](s []T) *Set[T] {
	m := make(map[T]struct{})
	for _, e := range s {
		m[e] = struct{}{}
	}
	es := make([]T, len(m))
	i := 0
	for k := range m {
		es[i] = k
		i++
	}
	return &Set[T]{es: es}
}

func (s *Set[T]) Elements() []T {
	return s.es
}

func LazySetIterator[T comparable](c []T) iter.Seq[T] {
	return func(yield func(v T) bool) {
		i := -1
		m := make(map[T]struct{})
		for {
			i++
			if len(c) <= i {
				return
			}
			_, ok := m[c[i]]
			if ok {
				continue
			}
			m[c[i]] = struct{}{}
			anotherValue := yield(c[i])
			if !anotherValue {
				return
			}
		}
	}
}

func lotsOfItems() []int {
	i := func(yield func(int) bool) {
		for i := 0; i < 100000000; i++ {
			yield(int(rand.Int31n(11)))
		}
	}
	items := make([]int, 0)
	for e := range i {
		items = append(items, e)

	}
	return items
}

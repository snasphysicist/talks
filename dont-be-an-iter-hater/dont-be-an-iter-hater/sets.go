package main

import (
	"fmt"
	"iter"
	"math/rand"
)

func setsMain() {
	// smallItems := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 9, 10, 8, 8, 5, 6, 7}
	fmt.Println("Full Collection")
	items := lotsOfItems()
	fmt.Println(len(items))
	fmt.Println("Eagerly Resolved")
	fmt.Println(len(NewSet(items).Elements()))
	fmt.Println("Lazily Resolved, Custom Iterator")
	i := IterateOver(NewSet(items))
	for p := 0; p < 5; p++ {
		n, ok := i.Next()
		if !ok {
			break
		}
		fmt.Println(*n)
	}
	fmt.Println("Lazily Resolved, Standard Iterator")
	count := 0
	for e := range NewSet(items).Iterator() {
		fmt.Println(e)
		count++
		if count == 5 {
			break
		}
	}
}

type Set[E comparable] struct {
	c []E
}

func NewSet[E comparable](c []E) *Set[E] {
	return &Set[E]{c: c}
}

// eagerly build set

func (s *Set[E]) Elements() []E {
	m := make(map[E]struct{})
	for _, e := range s.c {
		m[e] = struct{}{}
	}
	unique := make([]E, 0)
	for k := range m {
		unique = append(unique, k)
	}
	return unique
}

// lazily build set

type SetIterator[E comparable] struct {
	s        *Set[E]
	returned map[E]struct{}
	current  int
}

func IterateOver[E comparable](s *Set[E]) *SetIterator[E] {
	return &SetIterator[E]{s: s, returned: map[E]struct{}{}, current: 0}
}

func (i *SetIterator[E]) Next() (*E, bool) {
	for {
		if i.current >= len(i.s.c) {
			return nil, false
		}
		next := i.s.c[i.current]
		i.current++
		if _, ok := i.returned[next]; !ok {
			i.returned[next] = struct{}{}
			return &next, true
		}
	}
}

// iter.Seq approach

func (s *Set[E]) Iterator() iter.Seq[E] {
	current := 0
	returned := make(map[E]struct{})
	return func(yield func(E) bool) {
		// Not safe for concurrent access!
		for {
			if current >= len(s.c) {
				return
			}
			next := s.c[current]
			current++
			if _, ok := returned[next]; ok {
				continue
			}
			returned[next] = struct{}{}
			if !yield(next) {
				return
			}
		}
	}
}

// lots of items!

func lotsOfItems() []int {
	i := func(yield func(int) bool) {
		for i := 0; i < 100000000; i++ {
			yield(int(rand.Int31n(1000)))
		}
	}
	items := make([]int, 0)
	for e := range i {
		items = append(items, e)

	}
	return items
}

package main

import (
	"math/rand"
)

func SetsMain() {
	// smallSlice := []int{1, 2, 3, 4, 5, 4, 6, 2, 3, 10, 4, 6, 7, 8, 8, 9}
	// println("SMALL SLICE")
	// for _, e := range NewSet(smallSlice).Elements() {
	// 	println("Eager", e)
	// }
	// for e := range LazySetIterator(smallSlice) {
	// 	println("Lazy", e)
	// }
	// println("MAKING BIG SLICE")
	// bigSlice := LotsOfItems()
	// println("MADE BIG SLICE")
	// println("START RANGE-ING EAGER")
	// for _, e := range NewSet(bigSlice).Elements() {
	// 	println("Eager", e)
	// }
	// println("START RANGE-ING LAZY")
	// for e := range LazySetIterator(bigSlice) {
	// 	println("Lazy", e)
	// }
}

func LotsOfItems() []int {
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

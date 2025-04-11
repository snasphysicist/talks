package fallback

import "iter"

func rangeOverSliceOneArityMain() {
	for i := range OneArgumentSliceIterator([]string{"Amos", "Naomi", "Alex", "Holden"}) {
		println(i)
	}
}

func OneArgumentSliceIterator[T any](c []T) iter.Seq[int] {
	return func(yield func(int) bool) {
		i := 0
		for {
			if i >= len(c) {
				return
			}
			anotherValue := yield(i)
			if !anotherValue {
				return
			}
			i++
		}
	}
}

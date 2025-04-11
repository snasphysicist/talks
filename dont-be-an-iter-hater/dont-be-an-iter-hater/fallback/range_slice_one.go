package fallback

import "iter"

func rangeOverSliceOneArityMain() {
	for i := range OneArgumentSliceIterator([]string{"Amos", "Naomi", "Alex", "Holden"}) {
		println(i)
	}
}

func OneArgumentSliceIterator[T any](c []T) iter.Seq[int] {
	i := 0
	return func(yield func(int) bool) {
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

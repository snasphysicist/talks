package fallback

import "iter"

func rangeOverSliceTwoArityMain() {
	for i, v := range TwoArgumentSliceIterator([]string{"Amos", "Naomi", "Alex", "Holden"}) {
		println(i, v)
	}
}

func TwoArgumentSliceIterator[T any](c []T) iter.Seq2[int, T] {
	return func(yield func(k int, v T) bool) {
		i := 0
		for {
			if len(c) <= i {
				return
			}
			anotherValue := yield(i, c[i])
			if !anotherValue {
				return
			}
			i++
		}
	}
}

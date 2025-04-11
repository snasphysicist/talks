package fallback

import "iter"

func rangeOverSliceTwoArityMain() {
	for i, v := range TwoArgumentSliceIterator([]string{"Amos", "Naomi", "Alex", "Holden"}) {
		println(i, v)
	}
}

func TwoArgumentSliceIterator[T any](c []T) iter.Seq2[int, T] {
	i := 0
	return func(yield func(k int, v T) bool) {
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

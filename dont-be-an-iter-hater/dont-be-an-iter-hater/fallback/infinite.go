package fallback

import (
	"fmt"
	"iter"
	"math/rand"
)

func infiniteMain() {
	fmt.Println("Go Approach        ", goApproach())
	fmt.Println("Functional Approach", functionalApproach())
}

func goApproach() string {
	chars := []rune("ABCDEFG0123456789")
	id := ""
	for i := 0; i < 16; i++ {
		c := chars[rand.Int31n(16)]
		id = id + string(c)
	}
	return id
}

func functionalApproach() string {
	chars := []rune("ABCDEFG0123456789")
	rc := func() string {
		return string(chars[rand.Int31n(16)])
	}
	r := repeatedly(rc)
	t := take(16, r)
	id := join(t)
	return id
}

func repeatedly[T any](f func() T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for {
			if !yield(f()) {
				return
			}
		}
	}
}

func take[T any](n int, it iter.Seq[T]) iter.Seq[T] {
	count := 0
	return func(yield func(T) bool) {
		for e := range it {
			if count >= n {
				return
			}
			count++
			if !yield(e) {
				return
			}
		}
	}
}

func join[T any](it iter.Seq[T]) string {
	s := ""
	for e := range it {
		s = s + fmt.Sprint(e)
	}
	return s
}

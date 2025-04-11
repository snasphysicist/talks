package fallback

import (
	"iter"
	"reflect"
)

func rangeMapMain() {
	m := map[string]bool{
		"Rocinante":  false,
		"Canterbury": true,
		"Donnager":   true,
		"Razorback":  false,
	}
	for k, v := range mapIterator(m) {
		println("Ship", k, "was destroyed?", v)
	}
}

func mapIterator[T any](m map[string]T) iter.Seq2[string, T] {
	return func(yield func(string, T) bool) {
		ks := keys(m)
		i := 0
		for {
			if i >= len(ks) {
				return
			}
			anotherValue := yield(ks[i], m[ks[i]])
			if !anotherValue {
				return
			}
			i++
		}
	}
}

func keys[T any](m map[string]T) []string {
	vs := reflect.ValueOf(m).MapKeys()
	i := 0
	ks := make([]string, 0)
	for {
		if i >= len(vs) {
			return ks
		}
		ks = append(ks, vs[i].String())
		i++
	}
}

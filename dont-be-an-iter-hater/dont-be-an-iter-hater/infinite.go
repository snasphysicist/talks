package main

import (
	"fmt"
	"math/rand"
)

func InfiniteMain() {
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
	// chars := []rune("ABCDEFG0123456789")
	// rc := func() string {
	// 	return string(chars[rand.Int31n(16)])
	// }
	// r := repeatedly(rc)
	// t := take(16, r)
	// id := join(t)
	// return id
	return ""
}

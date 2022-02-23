package main

import (
	"fmt"
)

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	fmt.Println(60 / i)
}

func main() {
	for _, v := range []int{1, 2, 0, 4} {
		div60(v)
	}
}

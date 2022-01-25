package main

import (
	"fmt"
)

func main() {
	// NOTE: Using make you can specify type, length and optionally the capacity.
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x))
	y := make([]int, 5, 10)
	fmt.Println(y, len(y), cap(y))
	z := make([]int, 0, 10)
	fmt.Println(z, len(z), cap(z))
	z = append(z, 1, 2, 3, 4)
	fmt.Println(z, len(z), cap(z))
	var a []int     // nil slice
	var b = []int{} // non-nil slice
	fmt.Println(a, b)
	data := []int{2, 4, 6, 8}
	fmt.Println(data)
}

package main

import (
	"fmt"
)

func main() {
	// NOTE: if you need a slice that is independent of the original, use the
	// built-in copy function.
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x) // copies the entire slice, because they have the same size
	fmt.Println("num:", num)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	z := make([]int, 2)
	copy(z, x) // copies only the first 2 elements
	fmt.Println("z:", z)
	a := make([]int, 2)
	copy(a, x[2:]) // copies 2 element from the middle of x
	fmt.Println("a:", a)

	copy(x[:3], x[1:]) // overlapping elements will be overwritten
	fmt.Println(x)

	b := []int{1, 2, 3, 4}
	c := [4]int{4, 5, 6, 7}
	d := make([]int, 2)
	copy(d, c[:])
	fmt.Println(d)
	copy(c[:], b)
	fmt.Println(c)
}

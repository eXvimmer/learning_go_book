package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:] // not a copy, just a slice of a slice; i.e. shared memory

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	x[1] = 99 // changed everywhere
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// slices with overlapping storage
	m := []int{1, 2, 3, 4}
	n := m[:2]
	o := m[1:]
	m[1] = 20
	n[0] = 10
	o[1] = 30
	fmt.Println("n: ", n)
	fmt.Println("o: ", o)
	fmt.Println("m: ", m)

	// append makes overlapping slices more confusing
	r := []int{1, 2, 3, 4}
	s := r[:2]
	// NOTE:
	// cap(s) = cap(r) - the offset of the subslice within the original slice.
	// This means any unused capacity in the original slice is also shared with
	// any subslices.
	fmt.Println(cap(r), cap(s))
	s = append(s, 30)
	fmt.Println("r:", r)
	fmt.Println("s:", s)

	// Even more confusing slices
	x1 := make([]int, 0, 5)
	x1 = append(x1, 1, 2, 3, 4)
	y1 := x1[:2]
	z1 := x1[2:]
	fmt.Println(cap(x1), cap(y1), cap(z1)) // Read the NOTE above
	y1 = append(y1, 30, 40, 50)
	fmt.Println("y1:", y1, len(y1), cap(y1))
	fmt.Println("x1:", x1)
	x1 = append(x1, 60)
	fmt.Println("x1:", x1, len(x1), cap(x1))
	z1 = append(z1, 70)
	fmt.Println("z1:", z1, len(z1), cap(z1))

	// NOTE: Caution: either never use append with a subslice or make sure an
	// append doesn't cause an overwrite by using a full slice expression.
	// NOTE: The full slice expression protects against append. The full slice
	// expression includes a third part, which indicates the last position in the
	// parent slice's capacity that's available for subslice.
	fullSlice1 := x1[:2:2]  // cap = 2 (the last number) - offset
	fullSlice2 := x1[2:4:4] // cap = 4 - offset
	fmt.Println("fullSlice1:", fullSlice1, len(fullSlice1), cap(fullSlice1))
	fmt.Println("fullSlice2:", fullSlice2, len(fullSlice2), cap(fullSlice2))

}

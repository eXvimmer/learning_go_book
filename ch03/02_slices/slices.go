package main

import (
	"fmt"
)

func main() {
	var x = []int{1, 2, 3}
	var y = []int{1, 5: 6, 10: 100, 15}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x[0])
	var z []int // the zero value for this is nil
	// NOTE: nil means lack of value for some types. nil has no type, so it can
	// be assigned or compared against values of different types. A nil slice
	// contains nothing.
	// NOTE: Unlike arrays, slice isn't comparable (unless you use the reflect
	// package). You can only compare slices with nil.
	fmt.Println(z) // []
	fmt.Println(z == nil)
	fmt.Println(len(z))
	x = append(x, 4)
	fmt.Println(x)
	z = append(z, 1, 2, 3, 4)
	fmt.Println(z)
	z = append(z, x...)
	// ... operator expands the source slice into individual values
	fmt.Println(z)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 9, 13)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 99)
	fmt.Println(x, len(x), cap(x)) // double cap until 1024 and then add 25%
}

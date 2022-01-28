package main

import (
	"fmt"
)

// vals is variadic parameter
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))

	for _, v := range vals {
		out = append(out, base+v)
	}

	return out
}

func main() {
	// NOTE: The variadic parameter must be the last (or only) parameter in the
	// input parameter list. You indicate it with tree dots (...) before the
	// type.

	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))

	a := []int{4, 3}
	fmt.Println(addTo(3, a...)) // spreading a
	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))
}

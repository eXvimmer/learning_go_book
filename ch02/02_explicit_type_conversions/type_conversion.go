package main

import (
	"fmt"
)

func main() {
	var x int = 13
	var y float64 = 19.92
	var z float64 = float64(x) + y
	var d int = int(y) + x
	fmt.Println(z)
	fmt.Println(d)
	// NOTE: All type conversions in Go are explicit, because of this you cannot
	// treat another Go value as a boolean. Go doesn't allow truthiness. In fact,
	// no other type can be converted to a bool, implicitly or explicitly.
}

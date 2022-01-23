package main

import (
	"fmt"
)

const x int64 = 13

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// NOTE: error
	// x = x + 1
	// y = "bye"

	// fmt.Println(x)
	// fmt.Println(y)
}

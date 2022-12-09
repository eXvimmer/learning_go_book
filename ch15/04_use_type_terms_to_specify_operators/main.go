package main

import "fmt"

type BuiltInOrdered interface {
	int | int8 | int16 | int32 | int64 | string | float64 | float32 | uint |
		uint8 | uint16 | uint32 | uint64 | uintptr
}

func Min[T BuiltInOrdered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

type MyInt int

func main() {
	a, b := 10, 20
	fmt.Println(Min(a, b))

	// var (
	// 	MyA MyInt = 13
	// 	MyB MyInt = 99
	// )
	// fmt.Println(Min(MyA, MyB)) // ERROR
}

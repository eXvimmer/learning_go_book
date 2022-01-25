package main

import (
	"fmt"
)

func main() {
	var x [3]int
	var y = [3]int{1, 2, 3}
	// a sparse array
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	// [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 16]
	var a = [...]int{1, 2, 3}
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(a)
	fmt.Println(a == y)
	var d [2][3]int // multi-dimentional
	fmt.Println(d)
	fmt.Println(z[4], z[5])
	fmt.Println(len(x))
	// NOTE:
	// Go considers the size of the array to be part of the type of the array.
	// This makes an array that's declared to be [3]int a different type from an
	// array that's declared to be [4]int. This also means that you cannot use a
	// variable to specify the size of an array, because types must be resolved
	// at compile time, not at runtime. What's more, you can't use a type
	// conversion to
	// convert arrays of different sizes to identical types. Because you can't
	// convert arrays of different sizes into each other, you can't write a
	// function that works with arrays of any size and you can't assign arrays of
	// different sizes to the same variable.
}

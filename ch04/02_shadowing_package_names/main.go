package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Println(x)
	fmt := "oops"    // shadowning package names
	fmt.Println(fmt) // Error
}

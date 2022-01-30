package main

import (
	"fmt"
)

func failedUpdate(px *int) {
	x := 9
	px = &x
}

func update(px *int) {
	*px = 13
}

func main() {
	// NOTE: Another implication of copying a pointer is that if you want the
	// value assigned to a pointer parameter to still be there when you exit the
	// function, you must dereference the pointer and set the value. If you
	// change the pointer, you have changed the copy, not the original.
	// Dereferencing puts the new value in the memory location pointed to by both
	// the original and the copy.
	x := 4
	failedUpdate(&x) // copy the address of x into the parameter x
	fmt.Println(x)   // prints 4
	update(&x)
	fmt.Println(x) // prints 13
}

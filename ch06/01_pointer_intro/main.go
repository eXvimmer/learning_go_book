package main

import (
	"fmt"
)

type Foo struct{}

type Person struct {
	firstName  string
	middleName *string
	lastName   string
}

func main() {
	// NOTE: The zero value for a pointer is nil. Slices, maps, channels,
	// interfaces and functions are implemented with with pointers, so the zero
	// value for all of them is nil. nil is an untyped identifier that represents
	// the lack of a value for certain types. nil is a value defined in the
	// universe block and it can be shadowed.

	// NOTE: You cannot use & before primitive literal or a constant, because
	// they don't have memory addrsses.
	var x int32 = 13
	var y bool = true
	pointerX := &x
	pointerY := &y
	var pointerZ *string
	fmt.Println("x, pointerX, and *pointerX:", x, pointerX, *pointerX)
	fmt.Println("y, pointerY, and *pointerY", x, pointerY, *pointerY)
	fmt.Println("pointerZ", pointerZ)
	// fmt.Println("pointerZ, and *pointerZ", pointerZ, *pointerZ) // NOTE: panic
	// the write way to do this is to check for nil pointers
	if pointerZ != nil {
		fmt.Println("pointerZ, and *pointerZ", pointerZ, *pointerZ)
	}
	z := 86 + *pointerX
	fmt.Println(z)
	var pti *int32
	pti = &x
	fmt.Println(pti == pointerX)

	// NOTE: The built-in function new creates a pointer variable. It returns a
	// pointer to a zero value instance of the provided type.
	var n = new(int)      // n is *int with value of 0, not a nil pointer
	fmt.Println(n == nil) // false
	fmt.Println(*n)       // 0

	// for structs, use an & before struct literal to create a pointer instance.
	f := &Foo{}
	fmt.Println("f & *f:", f, *f)

	p := Person{
		firstName: "Mustafa",
		// middleName: "Amadeus Dovahkiin", // NOTE: panic, can't use a string for *string
		lastName: "Hayati",
	}
	fmt.Println(p, &p)
}

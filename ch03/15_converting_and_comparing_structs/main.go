package main

import (
	"fmt"
)

// NOTE: Go does allow you to perform a type conversion from one struct type
// to another if the fields of both structs have the same names, order, and
// types.

type firstPerson struct {
	name string
	age  int
}

type secondPerson struct {
	name string
	age  int
}

// NOTE: we can convert an instance of first person to secondPerson, but we
// cannot use == to compare an instance of firstPerson and an instance of
// secondPerson, becase they are different types.

type thirdPreson struct {
	age  int
	name string
}

// NOTE: we cannot convert an instance of firstPerson to thirdPerson, because
// fields are in a different order.

type fourthPreson struct {
	firstName string
	age       int
}

// NOTE: we cannot convert an instance of firstPerson to fourthPerson, because
// the field names don't match.

type fifthPerson struct {
	name          string
	age           int
	favoriteColor string
}

// NOTE: we cannot convert an instance of firstPerson to fifthPerson, because
// there's an additional field.

func main() {
	// NOTE: Structs that are entirely composed of comparable types are
	// comparable; those with slice or map fields are not, (function and channel
	// fields also prevent a struct from being comparable).

	f := firstPerson{
		name: "Mustafa",
		age:  29,
	}

	var g struct {
		name string
		age  int
	}
	// NOTE: You can also assign between named and anonymous  struct types if the
	// fields of both structs have have the same name, order and types.
	g = f
	fmt.Println(f == g)
}

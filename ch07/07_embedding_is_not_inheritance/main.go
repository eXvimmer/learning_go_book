package main

import (
	"fmt"
)

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

func main() {
	// NOTE: If you have a method on an embedded field that calls another method
	// on the embedded field, and the containing struct has a method of the same
	// name, the method on the embedded field will not invoke the method on the
	// containing struct.

	o := Outer{
		Inner: Inner{
			A: 13,
		},
		S: "Mustafa",
	}

	fmt.Println(o.Double()) // Inner: 26
}

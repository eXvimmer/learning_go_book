package main

import (
	"fmt"
)

func filedUpdate(g *int) {
	x := 10
	g = &x
}

func main() {
	// NOTE: When you pass a nil pointer to a function, you cannot make the value
	// non-nil. You can only reassign the value if there was a value already
	// assigned to the pointer. While confusing at first, it makes sense. Since
	// THE MEMORY LOCATION WAS PASSED TO THE FUNCTION VIA CALL-BY-VALUE, we can't
	// change the memory address, any more than we could change the value of an
	// int parameter. TODO: study figure 6-3

	var f *int // nil pointer
	filedUpdate(f)
	// NOTE: We COPY the value of f, which is nil, into the parameter g. This
	// means that g is also set to nil.
	fmt.Println(f)
}

package main

import "fmt"

/*
  extern int add(int a, int b);
*/
import "C"

//export doubler
func doubler(i int) int {
	return i * 2
}

// NOTE: If you do this, you can no longer declare C code directly in the
// comment before the import "C" statement. You can only declare functions, not
// define them; You then have to place your C code into a .c file in the same
// directory as your Go code and include the magic header "_cgo_export.h"

// WARN: Go is a garbage-collected language and C is not. This makes it
// difficult to integrate non- trivial Go code with C. While you can pass a
// pointer into C code, you cannot pass something that contains a pointer. This
// is very limiting, as things like strings, slices, and functions are
// implemented with pointers and therefore cannot be contained in a struct
// passed into a C function. That’s not all: a C function cannot store a copy
// of a Go pointer that lasts after the function returns. If you break these
// rules, your program will compile and run, but it may crash or behave
// incorrectly at runtime when the memory pointed to by the pointer is garbage
// collected. There are other limitations. For example, you cannot use cgo to
// call a variadic C function (such as printf). Union types in C are converted
// into byte arrays. And you cannot invoke a C function pointer (but you can
// assign it to a Go variable and pass it back into a C function).

func main() {
	sum := C.add(4, 9) // 4 * 2 + 9 = 17
	fmt.Println(sum)
}

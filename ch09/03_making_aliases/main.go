package main

import "fmt"

type Foo struct {
	x int
	S string
}

func (f Foo) hello() string {
	return "Hello"
}

func (f Foo) goodbye() string {
	return "Goodbye"
}

// NOTE: Bar is an alias which has the same fields and methods as the original
// type (Foo).
type Bar = Foo

// NOTE: alias is just another name for a type. If you want to add new methods
// or change the fields in an aliased struct, you must add them to the original
// type.

func MakeBar() Bar {
	bar := Bar{
		x: 13,
		S: "Mustafa",
	}
	var f Foo = bar
	fmt.Println(f.hello())
	return bar
	// return f // is also fine
}

func main() {
	f := MakeBar()
	fmt.Println(f.S)
}

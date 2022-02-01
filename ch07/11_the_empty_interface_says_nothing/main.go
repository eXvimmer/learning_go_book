package main

import (
	"fmt"
)

func main() {
	// NOTE: Sometimes you need a way to say a variable could store a value of any
	// type. Go uses interfaces to do this.
	var i interface{} // any value whose type implements zero or more methods
	i = 20
	i = "hello"
	i = struct {
		firstName string
		age       int
	}{
		firstName: "Mustafa",
		age:       29,
	}
	fmt.Println(i)
}

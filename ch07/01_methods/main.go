package main

import (
	"fmt"
)

// NOTE: Go supports methods on user defined types

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d.", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{
		FirstName: "Mustafa",
		LastName:  "Hayati",
		Age:       29,
	}

	fmt.Println(p.String())

}

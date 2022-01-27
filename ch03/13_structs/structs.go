package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
	pet  string
}

func main() {
	var mustafa person
	fmt.Println(mustafa.name, mustafa.age, mustafa.pet, mustafa)

	malena := person{} // all of the fields initialized to their zero values
	fmt.Println(malena.name, malena.age, malena.pet, malena)

	matt := person{
		"Matthew",
		26,
		"Dido",
	}
	fmt.Println(matt.name, matt.age, matt.pet, matt)

	david := person{
		age:  30,
		name: "David",
	}
	fmt.Println(david.name, david.age, david.pet, david)
}

package main

import (
	"fmt"
)

func main() {
	// NOTE: this is an anonymous struct, a struct without name
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "Mustafa"
	person.age = 29
	person.pet = "Peepa"

	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Dido",
		kind: "Rabbit",
	}
	fmt.Println(pet)

	developer := struct {
		name                    string
		age                     int
		mainProgrammingLanguage string
	}{
		name:                    "Mustafa",
		age:                     29,
		mainProgrammingLanguage: "C++",
	}

	fmt.Println(developer)
}

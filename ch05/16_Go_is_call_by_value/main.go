package main

import (
	"fmt"
)

type Person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p Person) {
	i *= 2
	s = "something random"
	p.name = "Lord Voldemort"
}

func modMap(m map[int]string) {
	m[2] = "Hello"
	m[3] = "World"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	// NOTE: you can modify any element in the slice, but you can't lengthen the
	// slice. This is true for maps and slices that are passed directly into
	// functions as well as map and slice fields in structs.
	s = append(s, 13)
}

func main() {
	// NOTE: when you supply a variable for a parameter to a function, Go ALWAYS
	// and ALWAYS makes a copy of the value of the variable. The behavior is a
	// little different for slices and maps. It's because maps and slices are
	// both implemented with pointers
	p := Person{
		name: "Mustafa",
		age:  29,
	}
	i := 13
	s := "Hacking VS Programming"
	m := map[int]string{
		1: "first",
		2: "second",
	}

	sl := []int{1, 2, 3}

	fmt.Println("Before modifyFails")
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	fmt.Println("-------------------")
	fmt.Println("After modifyFails")
	fmt.Println(i, s, p)
	fmt.Println("-------------------")
	fmt.Println("m Before modMap")
	fmt.Println(m)
	modMap(m)
	fmt.Println("-------------------")
	fmt.Println("m After modMap")
	fmt.Println(m)
	fmt.Println("-------------------")
	fmt.Println("sl Before modSlice")
	fmt.Println(sl)
	modSlice(sl)
	fmt.Println("-------------------")
	fmt.Println("sl After modSlice")
	fmt.Println(sl)
}

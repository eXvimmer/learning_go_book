package main

import (
	"fmt"
)

func main() {
	// NOTE: Go has 4 kinds of loops
	// A complete C-style for
	// A condition-only for
	// An infinite for
	// for-range

	// A complete C-style for
	for i := 0; i < 3; i++ { // you should use :=, var is not legal here
		fmt.Println(i)
	}

	// A condition-only for
	i := 3
	for i < 6 {
		fmt.Println(i)
		i += 1 // i++
	}

	// An infinite for
	i = 6
	for {
		fmt.Println(i)
		if i < 10 {
			break
		}
	}

	fmt.Println("----------------------------------")
	// for-range
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	fmt.Println("----------------------------------")
	oddVals := []int{1, 3, 5, 7, 9}
	// ignoring keys, using values
	for _, v := range oddVals {
		fmt.Println(v)
	}
	fmt.Println("----------------------------------")
	uniqueNames := map[string]bool{ // a fake set
		"Mustafa": true,
		"Malena":  true,
		"Maya":    true,
		"Mia":     true,
		"Melina":  true,
	}

	// using keys, ignoring values
	for k := range uniqueNames {
		fmt.Println(k)
	}
	fmt.Println("----------------------------------")

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// NOTE: don't expect any order when looping through maps
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
	fmt.Println("----------------------------------")

	// iterating over strings
	samples := []string{"Hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample { // r is a rune
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	// NOTE: strings are made out of bytes. When we iterate over strings with
	// for-range loop, it iterates over the runes, not the bytes. see p.75

	fmt.Println("----------------------------------")
	// NOTE: the for-range value is a copy
	for _, v := range evenVals {
		v *= 2 // v is a copy
	}
	fmt.Println(evenVals) // evenVals is not modified
}

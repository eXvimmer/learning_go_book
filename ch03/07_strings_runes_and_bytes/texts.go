package main

import (
	"fmt"
)

func main() {
	var s string = "Mustafa Hayati"
	var b byte = s[0]
	fmt.Println(s)
	fmt.Println(b)
	var surname string = s[8:]
	fmt.Println(surname)
	var s2 string = "Hello 🌞" // len = 10
	fmt.Println(len(s2))
	fmt.Println(s2)
	var s3 string = s2[4:7]
	var s4 string = s2[:5]
	var s5 string = s2[6:]
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	var a rune = 'x'
	var s6 string = string(a)
	var b1 byte = 'y'
	var s7 string = string(b1)
	fmt.Println(a, s6, b1, s7)
}


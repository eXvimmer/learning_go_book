package main

import (
	"fmt"
	"io"
)

type MyInt int

func doThings(i interface{}) {
	// NOTE: when an interface could be one of possible multiple types, use a
	// type switch

	// switch i := i.(type) { // NOTE: in this case shadowing is probably good
	switch j := i.(type) { // Type Switch
	case nil:
		fmt.Println(j, "i is nil, type of j is interface{}")
	case int:
		fmt.Println(j, "j is of type int")
	case MyInt:
		fmt.Println(j, "j is of type MyInt")
	case io.Reader:
		fmt.Println(j, "j is of type io.Reader")
	case string:
		fmt.Println(j, "j is of type string")
	case bool, rune:
		fmt.Println(j, "j is either a bool or rune, so j is of type interface{}")
	default:
		fmt.Println(j, "no idea what i is, so j is of type interface{}")
	}
}

func main() {
	var i interface{} // any
	var mine MyInt = 13
	i = mine
	i2 := i.(MyInt) // type assertion
	fmt.Println(i2 + 86)

	// i3 := i.(string) // panic, wrong type assertion
	// fmt.Println(i3)

	// i3 := i.(int) // panic, wrong type assertion, MyInt is not int
	// fmt.Println(i3 + 1)

	i3, ok := i.(int)
	if !ok {
		fmt.Printf("Unexpected type for %v", i)
	} else {
		fmt.Println(i3 + 1)
	}

	// NOTE: Type assertion and type conversion are very different. Type asserions
	// can only be applied to interfaces types and are checked at runtime.

}

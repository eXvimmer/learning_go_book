package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []string{"a", "b", "c"}
	sv := reflect.ValueOf(s)
	// NOTE: The Interface method returns the value of the variable as an empty
	// interface. However, the type information is lost; when you put the value
	// returned by Interface into a variable, you have to use type assertion to
	// get back to the right type
	s2, ok := sv.Interface().([]string)
	if !ok {
		panic("value of s is not of type []string")
	}
	fmt.Println(s2)

	// We can use reflection to set the value of a variable as well, but it's a
	// three-step process. First, you pass a pointer to the variable into
	// reflect.ValueOf. This returns a reflect.Value that represents the pointer.
	// The reason you need to pass a pointer to reflect.ValueOf to change the
	// value of the input parameter is that it is just like any other function in
	// Go (passed by value vs passed by pointer).Next, you need to get to the
	// actual value to set it. Finally, you get to the actual method that's used
	// to set the value. Just like there are special-case methods for reading
	// primitive types, there are special-case methods for setting primitive
	// types: SetBool, SetInt, SetFloat, SetString, and SetUint.
	i := 10
	iv := reflect.ValueOf(&i)
	ivv := iv.Elem()
	ivv.SetInt(13)
	fmt.Println(i)
}

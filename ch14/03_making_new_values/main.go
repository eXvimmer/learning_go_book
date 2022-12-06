package main

import (
	"fmt"
	"reflect"
)

var (
	stringType      = reflect.TypeOf((*string)(nil)).Elem() // convert nil to *string
	stringSliceType = reflect.TypeOf([]string(nil))         // nil to []string
)

func main() {
	// reflect.New creates a pointer to a scalar type, you can also use
	// reflection to do the same thing as the make keyword with the following
	// functions: MakeSlice, MakeChan, MakeMap, MakeMapWithSize

	sv := reflect.New(stringType).Elem() // similiar to the new function
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)

	sv.SetString("Mustafa")
	fmt.Println(sv)

	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)
}

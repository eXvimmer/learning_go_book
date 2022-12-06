package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int    `myTag:"value"`
	B string `myTag:"value2"`
}

func main() {
	var x int
	xt := reflect.TypeOf(x)
	fmt.Println(xt.Name()) // int

	xpt := reflect.TypeOf(&x)
	// NOTE: some types like slices or pointers, don't have names; in those
	// cases, Name returns an empty string.
	fmt.Println(xpt.Name())        // ""
	fmt.Println(xpt.Kind())        // reflect.Ptr
	fmt.Println(xpt.Elem().Name()) // int
	fmt.Println(xpt.Elem().Kind()) // reflect.Int

	fmt.Println("=======================================")
	f := Foo{}
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name()) // Foo
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	}
}

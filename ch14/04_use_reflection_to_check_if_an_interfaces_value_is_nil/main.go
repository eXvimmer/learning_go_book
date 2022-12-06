package main

import (
	"fmt"
	"reflect"
)

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() { // does it hold anything other than nil?
		return true
	}

	switch iv.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Func, reflect.Slice, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}

func main() {
	// NOTE: In order for an interface to be considered nil both the type and the
	// value must be nil.

	var a interface{}
	fmt.Println(a == nil, hasNoValue(a)) // true true

	var b *int
	fmt.Println(a == nil, hasNoValue(b)) // true true

	var c interface{} = b
	fmt.Println(c == nil, hasNoValue(c)) // false, true

	var d int
	fmt.Println(hasNoValue(d)) // false

	var e interface{} = d
	fmt.Println(e == nil, hasNoValue(e)) // false false
}

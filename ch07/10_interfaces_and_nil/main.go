package main

import (
	"fmt"
)

// NOTE: In order for an interface considered to be nil both the type and value
// must be nil.

func main() {
	var s *string
	var i interface{}     // type and value is nil
	fmt.Println(s == nil) // true
	fmt.Println(i == nil) // true
	i = s
	fmt.Println(i == nil) // false, type is not nil
	// NOTE: What nil indicates for an interface is whether or not you can invoke
	// methods on it. If an interface is non-nil, you can invoke methods on it.
	// It's not straight-forward to tell whether or not the value associated with
	// the interface is nil when the type is non-nil. You must use reflection to
	// find out.
}

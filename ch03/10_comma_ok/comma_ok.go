package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	if ok {
		fmt.Println(v, ok)
	}
	v, ok = m["world"] // NOTE: no colon(:) before =
	if ok {
		fmt.Println(v, ok)
	}
	v, ok = m["goodbye"]
	if ok {
		fmt.Println(v, ok)
	}
	fmt.Println(v, ok)
}

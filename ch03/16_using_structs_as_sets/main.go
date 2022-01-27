package main

import (
	"fmt"
)

// NOTE: Some people prefer to use struct{} for the value when a map is being
// used to implement a set. The advantage is that an empty struct uses zero
// bytes, while a boolean uses one byte. The disadvantage is that using a
// struct{} makes your code more clumsy. You have a less obvious assignment,
// and you need to use the comma ok idiom to check if a value is in the set:

func main() {
	intSet := map[int]struct{}{} // less memory
	vals := []int{5, 9, 3, 4, 0, 1, 13, 5, 9, 3, 3, 20, 99, 71, 77, 100}

	for _, v := range vals {
		intSet[v] = struct{}{} // an empty struct uses zero bytes
	}

	if _, ok := intSet[13]; ok { // but verbose
		fmt.Println("13 is in the set")
	}
}

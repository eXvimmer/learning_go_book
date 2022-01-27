package main

import (
	"fmt"
)

func main() {
	intSet := map[int]bool{} // fake set
	fmt.Println(intSet)
	vals := []int{5, 10, 2, 8, 2, 13, 9, 4, 4, 8, 99, 10, 0, 1, 3, 13}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(intSet)
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[13])
	fmt.Println(intSet[1000])
	if intSet[99] {
		fmt.Println("99 is in the set")
	}
}

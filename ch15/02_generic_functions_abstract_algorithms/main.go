package main

import (
	"fmt"
	"strings"
)

func Map[T any, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	r := []T{}
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T, U any](s []T, initializer U, f func(U, T) U) U {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

func main() {
	words := []string{"One", "Potato", "Two", "Potato"}
	names := []string{"Mustafa", "Margot", "Maya", "Sydney", "Malena", "Sadie"}

	ma := Filter(names, func(w string) bool {
		return strings.HasPrefix(w, "M") && strings.HasSuffix(w, "a")
	})
	fmt.Println(names)
	fmt.Println(ma)
	fmt.Println()

	filtered := Filter(words, func(w string) bool {
		return w != "Potato"
	})
	fmt.Println(words)
	fmt.Println(filtered)
	fmt.Println()

	lengths := Map(filtered, func(w string) int {
		return len(w)
	})
	fmt.Println(filtered)
	fmt.Println(lengths)
	fmt.Println()

	sum := Reduce(lengths, 0, func(i1, i2 int) int {
		return i1 + i2
	})
	fmt.Println(lengths)
	fmt.Println(sum)
	fmt.Println()
}

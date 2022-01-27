package main

import (
	"fmt"
)

func main() {
	var nilMap map[string]int // the zero value for a map is nil
	fmt.Println(nilMap)
	// teams := map[string][]string{
	// 	"Orcas":   []string{"Fred", "Ralph", "Bijou"},
	// 	"Lions":   []string{"Sarah", "Peter", "Billie"},
	// 	"Kittens": []string{"Waldo", "Raul", "Ze"},
	// }
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)
	ages := make(map[string][]string, 10)
	fmt.Println(ages)
	fmt.Println(len(ages))

	totalWins := map[string]int{}
	fmt.Println(totalWins)
	totalWins["Orcas"] = 1
	fmt.Println(totalWins)
	totalWins["Lions"] = 2
	fmt.Println(totalWins)
	fmt.Println(totalWins["Kittens"]) // not created, zero value returned
	fmt.Println(totalWins)
	totalWins["Kittens"]++
	// now it's created. This works, because map returns its zero value by default
	fmt.Println(totalWins)
	totalWins["Lions"] = 3
	fmt.Println(totalWins)
}

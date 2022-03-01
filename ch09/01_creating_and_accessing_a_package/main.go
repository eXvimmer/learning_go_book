package main

import (
	"fmt"

	"github.com/exvimmer/package_example_one/formatter"
	"github.com/exvimmer/package_example_one/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}

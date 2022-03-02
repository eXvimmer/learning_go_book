package foo

import "github.com/exvimmer/the_internal_package_example/foo/internal"

func UseDoubler(i int) int {
	return internal.Doubler(i)
}

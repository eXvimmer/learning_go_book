package sibling

import "github.com/exvimmer/the_internal_package_example/foo/internal"

func AlsoUseDoubler(i int) int {
	return internal.Doubler(i)
}

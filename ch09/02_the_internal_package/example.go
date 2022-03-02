package internal_package_example

import "github.com/exvimmer/the_internal_package_example/foo/internal"

func FailUseDoubler(a int) int {
	return internal.Doubler(a)
}

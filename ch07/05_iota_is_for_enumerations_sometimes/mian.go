package main

import ()

// NOTE: When using iota, the best practice is to first define a type based on
// int that will represent all of the valid values:
type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal                   // same type as above, value incremented
	Spam
	Social
	Advertisement
)

// NOTE: when a new const block is created, iota is set back to 0.
// NOTE: Use iota for "internal" purpose only. That is, where constants are
// referred to by name rather than by value. If the actual value matters,
// specify it explicitly.

type BitField int

// WARN: Be careful when using this pattern. Document what you're doing. This
// pattern is fragile.
const (
	Field1 BitField = 1 << iota
	Field2
	Field3
	Field4
)

func main() {

}

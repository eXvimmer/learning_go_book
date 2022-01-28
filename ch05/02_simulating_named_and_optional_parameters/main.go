package main

import ()

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	// do something here
}

func main() {
	// NOTE: named and optional input parameters are not preset in Go. You must
	// supply all of the parameters for a function (with 1 exception).

	// NOTE: if you want to `emulate` named and optional parameters, define a
	// struct that has fields that match the desired parameters, and pass the
	// struct to your function.

	MyFunc(MyFuncOpts{
		FirstName: "Mustafa",
		Age:       29,
	})

	MyFunc(MyFuncOpts{
		FirstName: "Mustafa",
		LastName:  "Hayati",
	})
}

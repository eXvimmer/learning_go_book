package main

import ()

type Stringer interface { // interfaces are usually named with "er" endings.
	String() string
}

// NOTE: 🚀 In n interface declaration, an interface literal appears after the
// name of the interface type. It lists the methods that must be implemented by
// a concrete type to meet the interface. The methods defined by an interface
// are called the method set of the interface.a

// NOTE: Interfaces specify what callers need. The client code defines the
// interface to specify what functionality it requires.

func main() {
}

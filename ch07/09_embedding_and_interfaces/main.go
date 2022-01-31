package main

import ()

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type ReaderCloser interface {
	Reader
	Closer
}

func main() {

}

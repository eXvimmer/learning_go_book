package main

func example() {
	defer func() int {
		return 2 // NOTE: there's no way to read this value
	}()
}

func main() {
}

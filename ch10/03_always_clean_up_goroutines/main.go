package main

import "fmt"

func countTo(max int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func main() {
	// NOTE: when you use all of the values, the goroutine exits
	for i := range countTo(10) {
		fmt.Println(i)
	}

	fmt.Println("==============")
	// NOTE: if we exit the loop early, the goroutine blocks forever, waiting for
	// a value to be read from the channel
	for i := range countTo(10) {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

package main

import "fmt"

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		v := 1
		ch1 <- v // NOTE: this goroutine cannot proceed until ch1 is read
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v // NOTE: the main goroutine cannot proceed until ch2 is read
	v2 := <-ch1
	// NOTE: if we swap the 2 lines above, the code will work
	fmt.Println(v, v2)
}

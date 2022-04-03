package main

import "fmt"

func main() {
	a := []int{2, 4, 6, 8, 10}
	ch1 := make(chan int, len(a))

	// NOTE: this one is buggy: the reason why every goroutine wrote 20 to the
	// ch1 is that the closure for every goroutine captured the same variable.
	for _, v := range a {
		go func() {
			ch1 <- v * 2
		}()
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch1)
	}

	fmt.Println("*********************************")
	// solution #1
	for _, v := range a {
		v := v
		go func() {
			ch1 <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch1)
	}

	fmt.Println("*********************************")
	// solution #2
	for _, v := range a {
		go func(val int) {
			ch1 <- val * 2
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch1)
	}
	// NOTE: any time your goroutine usees a variable whose value might change,
	// pass the current value of the variable into the goroutine
}

package main

import "fmt"

/* NOTE:
 * You need to properly handle closed channels. If one of the cases in a select
 * is reading a closed channel, it will always be successful, returning the
 * zero value. Every time that case is selected, you need to check to make sure
 * that the value is valid and skip the case. If reads are spaced out, your
 * program is going to waste a lot of time reading junk values. When that
 * happens, we rely on something that looks like an error: reading a nil
 * channel. Reading from or writing to a nil channel causes your code to hang
 * forever. While that is bad if it is triggered by a bug, you can use a nil
 * channel to disable a case in a select. When you detect that a channel has
 * been closed, set the channel’s variable to nil. The associated case will no
 * longer run, because the read from the nil channel never returns a value:
 */

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan struct{})
	for {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil // this case will never succeed again
				continue
			}
			fmt.Println(v)
		case v, ok := <-ch2: // this case will never succeed again
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(v)
		case <-done:
			break
		}
	}
}

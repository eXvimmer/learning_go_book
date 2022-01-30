package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in updateWrong", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in udpateRight", c.String())
}

func main() {
	var c Counter
	fmt.Println(c.String())
	c.Increment() // NOTE: This is converted to (&c).Increment()
	// NOTE: when you use a local variable that's a value type, Go automatically
	// converts it to a point type.
	fmt.Println(c.String())

	doUpdateWrong(c)
	fmt.Println("in main", c.String())
	doUpdateRight(&c)
	fmt.Println("in main", c.String())
}

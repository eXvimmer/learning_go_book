package main

import (
	"fmt"
	"time"
)

func main() {
	// NOTE: Go doesn’t use yyyy-mm-dd layout to format a time. Instead, you
	// format a special layout parameter
	// Mon Jan 2 15:04:05 MST 2006
	// the same way as the time or date should be formatted. (This date is easier
	// to remember when written as 01/02 03:04:05PM ‘06 -0700.)

	/*
		Layout options
		Type	    Options
		Year	    06   2006
		Month	    01   1   Jan   January
		Day	      02   2   _2   (width two, right justified)
		Weekday 	Mon   Monday
		Hours   	03   3   15
		Minutes 	04   4
		Seconds  	05   5
		ms        μs ns	.000   .000000   .000000000
		ms        μs ns	.999   .999999   .999999999   (trailing zeros removed)
		am/pm	    PM   pm
		Timezone	MST
		Offset  	-0700   -07   -07:00   Z0700   Z07:00
	*/

	// string to time.Time
	t, err := time.Parse("2006-01-02 15:04:05 -0700", "2022-11-02 09:20:30 -0330")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// time.Time to string
	fmt.Println(t.Format("January 02 2006 03:04:05PM MST"))
	fmt.Println(t.Format("Jan 02 2006 15:04:05 -0700"))
	fmt.Println(t.Format("2006-01-02T15:04:05 MST"))
	fmt.Println(t.Format("January 2, 2006 at 03:04:05PM MST"))
	fmt.Println(t.Format("Mon January 02, 2006"))
	fmt.Println(t.Format("January 2, 2006 at 03:04:05PM Z07:00"))
	fmt.Println(t.Format("Monday, January 2 at 03PM MST"))
	fmt.Println(t.Format(time.RubyDate))
	fmt.Println(t.Format(time.Kitchen))
}

package main

import (
	"errors"
	"log"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{} // ch is a buffered channel of "tokens"
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch: // read token from channel
		f()
		pg.ch <- struct{}{} // return the token to the buffered channel
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)

	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}

	return &PressureGauge{
		ch: ch,
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(time.Second * 2)
	return "done"
}

func main() {
	// Another technique that can be implemented with a buffered channel is
	// backpressure. It is counterintuitive, but systems perform better overall
	// when their components limit the amount of work they are willing to
	// perform. We can use a buffered channel and a select statement to limit the
	// number of simultaneous requests in a system
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})

		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(http.StatusText(http.StatusTooManyRequests)))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

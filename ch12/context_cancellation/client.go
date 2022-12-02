package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

var client = http.Client{}

func callBoth(ctx context.Context, errVal, slowURL, fastURL string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	// NOTE: Any time you create a context that has an associated cancel
	// function, you must call that cancel function when you are done processing,
	// whether or not your processing ends in an error. If you do not, your
	// program will leak resources (memory and goroutines) and eventually slow
	// down or crash. There is no error if you call the cancel function more than
	// once; any invocation after the first does nothing.
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowURL)
		if err != nil {
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastURL+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()
	wg.Wait()
	fmt.Println("done with both")
}

func callServer(ctx context.Context, label, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "request err:", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "response err:", err)
		return err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(label, "read err:", err)
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result:", result)
	}
	if result == "error" {
		fmt.Println("cancelling from", label)
		return errors.New("error happened")
	}
	return nil
}

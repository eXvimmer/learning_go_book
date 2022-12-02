package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// NOTE: The net/http package defines a Client type to make HTTP requests and
	// receive HTTP responses. A default client instance (the DefaultClient) is
	// found in the net/http package, but you should avoid using it in production
	// applications, because it defaults to having no timeout. Instead,
	// instantiate your own. You only need to create a single http.Client for
	// your entire program, as it properly handles multiple simultaneous requests
	// across goroutines.

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// create a new *http.Request
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet,
		"https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		panic(err)
	}
	// set headers before send
	req.Header.Add("Something-Header-Related", "Mustafa Hayati")
	// send a request
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.StatusCode))
	}
	fmt.Println(res.Header.Get("Content-Type"))
	var data struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	// read data form response
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", data)
	// WARN: There are functions in the net/http package to make GET, HEAD, and
	// POST calls. Avoid using these functions because they use the default
	// client, which means they don't set a request timeout.
}

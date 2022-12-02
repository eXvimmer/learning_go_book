package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// NOTE: you can use json.Decoder and json.Encoder to read from or write to
	// multiple JSON structs.

	const data = `
    {"name": "Mustafa", "age": 30}
    {"name": "Malena", "age": 26}
    {"name": "Maya", "age": 24}
  `
	var (
		t Person
		// our final container
		b bytes.Buffer
	)
	dec := json.NewDecoder(strings.NewReader(data)) // read from data
	enc := json.NewEncoder(&b)                      // write to b

	for dec.More() { // read the data, one JSON object at a time
		err := dec.Decode(&t) // read to t
		if err != nil {
			panic(err)
		}
		fmt.Println(t)

		err = enc.Encode(t) // write from t to buffer
		if err != nil {
			panic(err)
		}
	}
	out := b.String()
	fmt.Println(out)
	// NOTE: Our example has multiple JSON objects in the data stream that aren't
	// wrapped in an array, but you can also use the json.Decoder to read a
	// single object from an array without loading the entire array into memory
	// at once. This can greatly increase performance and reduce memory usage.
}

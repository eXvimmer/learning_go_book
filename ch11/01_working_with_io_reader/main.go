package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

// counts English letters in anything that conforms to io.reader
func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		// NOTE: There is one unusual thing about the Read method in io.Reader.
		// In most cases when a function or method has an error return value, we
		// check the error before we try to process the nonerror return values.
		// We do the opposite for Read because there might have been bytes
		// returned before an error was triggered by the end of the data stream
		// or by an unexpected condition.
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func buildGzipReader(filename string) (*gzip.Reader, func(), error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(f)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		f.Close()
	}, nil
}

func main() {
	s := "the quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s) // create a new io.Reader
	m, err := countLetters(sr)
	if err != nil {
		panic(err)
	}
	fmt.Println(m)

	fmt.Println("====================================")
	r, closer, err := buildGzipReader("./my_data.txt.gz")
	if err != nil {
		panic(err)
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		panic(err)
	}
	fmt.Println(counts)
}

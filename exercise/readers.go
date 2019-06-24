package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
	var i int
	for i = range b {
		// filled by 'A'
		b[i] = 'A'
	}
	return i, nil
}

func main() {
	reader.Validate(MyReader{})
}

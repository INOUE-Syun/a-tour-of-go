package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r13.r.Read(p)
	for i, v := range p {
		// see https://en.wikipedia.org/wiki/ROT13
		switch {
		case v >= 'A' && v < 'N', v >= 'a' && v < 'n':
			p[i] += 13
		case v >= 'N' && v <= 'Z', v >= 'n' && v <= 'z':
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

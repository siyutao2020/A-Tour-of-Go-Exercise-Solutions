package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(bytes []byte) (int, error) {
	l, e := r13.r.Read(bytes)

	for i := range bytes {
		if bytes[i] <= 'M' || (bytes[i] >= 'a' && bytes[i] <= 'm') {
			bytes[i] = bytes[i] + 13
		} else {
			bytes[i] = bytes[i] - 13
		}

	}
	return l, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
	l, err := r13.r.Read(b)
	if err != nil {
		return l, err
	}
	for i, v := range b {
		var c byte
		switch {
		case v >= 'A' && v <= 'Z':
			c = ((v - 'A' + 13) % 26) + 'A'
		case v >= 'a' && v <= 'z':
			c = ((v - 'a' + 13) % 26) + 'a'
		default:
			continue
		}
		b[i] = c
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}


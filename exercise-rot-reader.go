package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	// find ASCII code offset and remove, rotate, reapply offset
	// More efficient than a lookup table
	if b >= 'A' && b <= 'Z' {
		b = 'A' + (b-'A'+13)%26
	} else if b >= 'a' && b <= 'z' {
		b = 'a' + (b-'a'+13)%26
	}
	return b
}

func (T rot13Reader) Read(b []byte) (int, error) {
	read, err := T.r.Read(b)
	if err != nil {
		return read, err
	}

	for i, v := range b {
		b[i] = rot13(v)
	}

	return read, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

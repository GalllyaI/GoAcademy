package main

import (
	"io"
	"os"
)

type ReverseStringReader struct {
	input string
}

func NewReverseStringReader(str string) *ReverseStringReader {
	return &ReverseStringReader{input: str}
}

func (r *ReverseStringReader) Read(buf []byte) (n int, err error) {

	revStr := ""
	for _, s := range r.input {
		revStr = string(s) + revStr
	}

	n = copy(buf, []byte(revStr))

	return n, io.EOF
}

func main() {
	in := "input"
	rr := NewReverseStringReader(in)
	io.Copy(os.Stdout, rr)

}

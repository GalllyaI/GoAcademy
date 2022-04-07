package main

import (
	"bytes"
	"io"
	"testing"
)

func TestReader(t *testing.T) {
	var buf bytes.Buffer
	rr := NewReverseStringReader("input")
	io.Copy(&buf, rr)
	expected := "tupni"

	if buf.String() != expected {
		t.Errorf("got %s, want %s", buf.String(), expected)
	}

}

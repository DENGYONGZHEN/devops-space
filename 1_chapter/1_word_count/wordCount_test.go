package main

import (
	"bytes"
	"testing"
)

func TestWordCount(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4
	res := count(b)
	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

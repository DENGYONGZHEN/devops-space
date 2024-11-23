package main

import (
	"bytes"
	"testing"
)

func TestCountWordFlag(t *testing.T) {

	b := bytes.NewBufferString("word1 word2 word3 word4\n")
	exp := 4

	res := countWordFlag(b, "word")

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLineFlag(t *testing.T) {

	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")
	exp := 3
	res := countWordFlag(b, "line")
	if exp != res {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func countWordFlag(r io.Reader, countType string) int {

	scanner := bufio.NewScanner(r)

	if countType == "word" {
		scanner.Split(bufio.ScanWords)
	}
	if countType == "byte" {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	//Defining a boolean flag -l to count lines instead of words
	countType := flag.String("c", "line", "Count lines")

	//Parsing the flags provided by the user
	flag.Parse()
	fmt.Println(countWordFlag(os.Stdin, *countType))
}

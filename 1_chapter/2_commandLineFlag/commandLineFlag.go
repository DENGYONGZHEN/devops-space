package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func countWordFlag(r io.Reader, countLine bool) int {

	scanner := bufio.NewScanner(r)

	if !countLine {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}
	return wc
}

func main() {
	//Defining a boolean flag -l to count lines instead of words
	countLine := flag.Bool("l", false, "Count lines")
	//Parsing the flags provided by the user
	flag.Parse()
	fmt.Println(countWordFlag(os.Stdin, *countLine))
}

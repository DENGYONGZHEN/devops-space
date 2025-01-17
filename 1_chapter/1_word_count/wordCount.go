package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	// A scanner is used to read text form a Reader (such as files)
	scanner := bufio.NewScanner(r)

	//Define the scanner split type to words(default is split by lines)
	scanner.Split(bufio.ScanWords)

	//Define a counter
	wc := 0

	//For every word scanned,increment the counter
	for scanner.Scan() {
		wc++
	}
	//Return the total
	return wc
}

func main() {
	fmt.Println(count(os.Stdin))
}

/*
Copyright © 2025 deng

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import "pScan/cmd"

func main() {
	cmd.Execute()
}

//  1 Allow the user to provide port ranges, such as 1-1024, in addition to specific
//  ports for scan.
//  2 Validate the provided port numbers are within the proper range for TCP
//  ports from 1 to 65535.
//  3 Allow the user to execute UDP port scans in addition to TCP. Update the
//  scan package and the command-line tool accordingly.
//  4 Add a new flag to the scan subcommand allowing the user to specify a filter
//  to show only open or closed ports.
//  5 Add a new flag to the scan subcommand allowing the user to specify a
//  custom timeout for the scan.

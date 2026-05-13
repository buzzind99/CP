//go:build codeforces_1741A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func abs(x int) int {
    if x < 0 { return -x }
    return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		a, b := next(), next()
		sizeA, sizeB := a[len(a)-1], b[len(b)-1]

		if sizeA == sizeB {
			if sizeA == byte('S') {
				if len(a) > len(b) {
					fmt.Fprintln(wr, "<")
				} else if len(a) < len(b) {
					fmt.Fprintln(wr, ">")
				} else {
					fmt.Fprintln(wr, "=")
				}
			} else if sizeA == byte('M') {
				fmt.Fprintln(wr, "=")
			} else {
				if len(a) > len(b) {
					fmt.Fprintln(wr, ">")
				} else if len(a) < len(b) {
					fmt.Fprintln(wr, "<")
				} else {
					fmt.Fprintln(wr, "=")
				}
			} 
		} else if sizeA > sizeB {
			fmt.Fprintln(wr, "<")
		} else {
			fmt.Fprintln(wr, ">")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1741/A
  Tags: implementation, strings
  Rating: 800
*/

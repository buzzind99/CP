//go:build codeforces_1850B

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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		maxVal, idx := -1, -1
		for i := range n {
			a, b := nextInt(), nextInt()
			if a <= 10 {
				if b > maxVal {
					maxVal = b
					idx = i+1
				}
			}
		}

		fmt.Fprintln(wr, idx)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1850/B
  Tags: implementation, sortings
  Rating: 800
*/

//go:build codeforces_2185B

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
		maxVal := -1
		for range n {
			maxVal = max(maxVal, nextInt())
		}

		fmt.Fprintln(wr, maxVal*n)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2185/B
  Tags: greedy
  Rating: 800
*/

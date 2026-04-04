//go:build codeforces_2218B

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
		maxVal := -68
		sum := 0
		for range 7 {
			a := nextInt()
			maxVal = max(maxVal, a)
			sum -= a
		}
		sum += 2*maxVal

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2218/B
  Tags: -
  Rating: -
  Contest: Codeforces Round 1090 (Div. 4)
*/

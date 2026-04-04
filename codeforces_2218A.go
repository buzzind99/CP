//go:build codeforces_2218A

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

		fmt.Fprintln(wr, n)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2218/A
  Tags: -
  Rating: -
  Contest: Codeforces Round 1090 (Div. 4)
*/

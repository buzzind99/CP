//go:build codeforces_2179A

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
		k, x := nextInt(), nextInt()

		fmt.Fprintln(wr, k*x+1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2179/A
  Tags: math, strings
  Rating: 800
*/

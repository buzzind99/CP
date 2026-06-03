//go:build codeforces_1768A

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
		k := nextInt()

		fmt.Fprintln(wr, k-1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1768/A
  Tags: greedy, math, number theory
  Rating: 800
*/

//go:build codeforces_2094B

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
		_, m, l, _ := nextInt(), nextInt(), nextInt(), nextInt()
		left := min(abs(l), m)
		right := m-left

		fmt.Fprintln(wr, -left, right)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2094/B
  Tags: brute force, constructive algorithms
  Rating: 800
*/

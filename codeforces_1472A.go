//go:build codeforces_1472A

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

func powInt(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		w, h, n := nextInt(), nextInt(), nextInt()
		count := 0
		for w%2 == 0 || h%2 == 0 {
			if w%2 == 0 { w /= 2; count++ }
			if h%2 == 0 { h /= 2; count++ }
			if powInt(2, count) >= n { break }
		}
		possible := "NO"
		if powInt(2, count) >= n { possible = "YES" }

		fmt.Fprintln(wr, possible)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1472/A
  Tags: greedy, math
  Rating: 800
*/

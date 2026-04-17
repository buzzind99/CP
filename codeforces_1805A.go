//go:build codeforces_1805A

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

func isEven (n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		sumXor := nextInt()
		for range n-1 {
			sumXor ^= nextInt()
		}

		if isEven(n) {
			if sumXor == 0 {
				fmt.Fprintln(wr, 0)
			} else {
				fmt.Fprintln(wr, -1)
			}
		} else {
			fmt.Fprintln(wr, sumXor)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1805/A
  Tags: bitmasks, brute force
  Rating: 800
*/

//go:build codeforces_2121A

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
		n, s := nextInt(), nextInt()
		maxVal, minVal := s, s
		for range n {
			x := nextInt()
			maxVal, minVal = max(maxVal, x), min(minVal, x)
		}

		distMax, distMin, maxRange := maxVal-s, s-minVal, maxVal-minVal
		sol := distMax+maxRange
		if distMax > distMin { sol = distMin+maxRange }

		fmt.Fprintln(wr, sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2121/A
  Tags: brute force, math
  Rating: 800
*/

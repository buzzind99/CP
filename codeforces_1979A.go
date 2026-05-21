//go:build codeforces_1979A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math"
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
		prev := nextInt()
		minVal := math.MaxInt
		for range n-1 {
			curr := nextInt()
			currMax := max(curr, prev)
			minVal = min(minVal, currMax)
			prev = curr
		}

		fmt.Fprintln(wr, minVal-1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1979/A
  Tags: brute force, greedy, implementation
  Rating: 800
*/

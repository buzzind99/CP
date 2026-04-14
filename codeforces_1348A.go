//go:build codeforces_1348A

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
		n := nextInt()
		sumA := 1 << n
		sumB := 0

		for i := 1; i < n/2; i++ {
			sumA += (1 << i)
		}
		for i := n/2; i < n; i++ {
			sumB += (1 << i)
		}

		fmt.Fprintln(wr, abs(sumA - sumB))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1348/A
  Tags: greedy, math
  Rating: 800
*/

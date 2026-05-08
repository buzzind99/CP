//go:build codeforces_2051A

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

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		a := make([]int, n)
		for i := range n {
			a[i] = nextInt()
		}
		b := make([]int, n)
		for i := range n {
			b[i] = nextInt()
		}

		diff := 0
		for i := 0; i < n-1; i++ {
			if a[i] > b[i+1] {
				diff += (a[i] - b[i+1])
			}
		}
		diff += a[n-1]

		fmt.Fprintln(wr, diff)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2051/A
  Tags: greedy
  Rating: 800
*/

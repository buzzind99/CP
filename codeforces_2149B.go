//go:build codeforces_2149B

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
		arr := make([]int, 0, n)
		for range n {
			a := nextInt()
			arr = append(arr, a)
		}
		slices.Sort(arr)

		maxDiff := -1
		for i := 0; i < n-1; i += 2 {
			maxDiff = max(maxDiff, arr[i+1]-arr[i])
		}

		fmt.Fprintln(wr, maxDiff)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2149/B
  Tags: greedy, sortings
  Rating: 800
*/

//go:build codeforces_2179C

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"slices"
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

func gcd(a, b int) int {
	for b != 0 { a, b = b, a%b }
	return a
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

		ans := arr[0]
		if n >= 2 {
			cand := arr[1]-arr[0]
			if cand > ans {
				ans = cand
			}
		}

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2179/C
  Tags: implementation, math, number theory, sortings
  Rating: 800
*/

//go:build codeforces_2114A

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

func getRoot(n int) int {
	low, high := 0, 2000000000
	ans := -1
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid <= n {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		root := getRoot(n)

		if root*root == n {
			fmt.Fprintln(wr, 0, root)
		} else {
			fmt.Fprintln(wr, -1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2114/A
  Tags: binary search, brute force, math
  Rating: 800
*/

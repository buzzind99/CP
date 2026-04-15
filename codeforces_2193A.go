//go:build codeforces_2193A

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
		n, s, x := nextInt(), nextInt(), nextInt()
		sum := 0
		for range n {
			sum += nextInt()
		}

		gap := s-sum
		if gap < 0 || gap%x != 0 {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2193/A
  Tags: brute force, math
  Rating: 800
*/

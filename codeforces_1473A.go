//go:build codeforces_1473A

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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, d := nextInt(), nextInt()
		arr := make([]int, 0, n)
		maxVal := -1
		for range n {
			a := nextInt()
			arr = append(arr, a)
			maxVal = max(maxVal, a)
		}

		if maxVal <= d {
			fmt.Fprintln(wr, "YES")
		} else {
			slices.Sort(arr)
			if arr[0]+arr[1] <= d {
				fmt.Fprintln(wr, "YES")
			} else {
				fmt.Fprintln(wr, "NO")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1473/A
  Tags: greedy, implementation, math, sortings
  Rating: 800
*/

//go:build codeforces_1980B

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
		n, f, k := nextInt(), nextInt(), nextInt()
		arr := make([]int, 0, n)
		for range n {
			arr = append(arr, nextInt())
		}
		fav := arr[f-1]
		slices.Sort(arr)
		slices.Reverse(arr)

		if arr[k-1] > fav {
			fmt.Fprintln(wr, "NO")
		} else if arr[k-1] < fav {
			fmt.Fprintln(wr, "YES")
		} else {
			if k < n && arr[k] == fav {
				fmt.Fprintln(wr, "MAYBE")
			} else {
				fmt.Fprintln(wr, "YES")
			}
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1980/B
  Tags: sortings
  Rating: 800
*/

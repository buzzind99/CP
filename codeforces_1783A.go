//go:build codeforces_1783A

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
		n := nextInt()
		arr := make([]int, 0, n)
		for range n {
			arr = append(arr, nextInt())
		}
		slices.Reverse(arr)

		if slices.Max(arr) == slices.Min(arr) {
			fmt.Fprintln(wr, "NO")
			continue
		}
		if arr[0] == arr[1] {
			arr[1], arr[n-1] = arr[n-1], arr[1]
		}
		fmt.Fprintln(wr, "YES")
		for _, v := range arr {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1783/A
  Tags: constructive algorithms, math, sortings
  Rating: 800
*/

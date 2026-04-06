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
		slices.Sort(arr)

		if arr[0] == arr[n-1] {
			fmt.Fprintln(wr, "NO")
			continue
		}
		fmt.Fprintln(wr, "YES")
		fmt.Fprint(wr, arr[n-1], " ")
		for i := range n - 1 {
			fmt.Fprint(wr, arr[i], " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1783/A
  Tags: constructive algorithms, math, sortings
  Rating: 800
*/

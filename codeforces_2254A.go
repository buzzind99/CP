//go:build codeforces_2254A

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
		a, b, c := nextInt(), nextInt(), nextInt()
		arr := []int{a,b,c}
		slices.Sort(arr)

		fmt.Fprintln(wr, min(arr[2]-arr[1], arr[1]-arr[0]))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2254/A
  Tags: implementation, sortings
  Rating: 800
*/

//go:build codeforces_1986A

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
		arr := []int{nextInt(), nextInt(), nextInt()}
		slices.Sort(arr)

		fmt.Fprintln(wr, arr[2]-arr[0])
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1986/A
  Tags: brute force, geometry, math, sortings
  Rating: 800
*/

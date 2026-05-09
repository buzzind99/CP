//go:build codeforces_1929A

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
			a := nextInt()
			arr = append(arr, a)
		}
		slices.Sort(arr)

		sum := 0
		for i := range n-1 {
			sum += arr[i+1]-arr[i]
		}

		fmt.Fprintln(wr, sum)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1929/A
  Tags: constructive algorithms, greedy, math, sortings
  Rating: 800
*/

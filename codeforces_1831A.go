//go:build codeforces_1831A

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

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()
		arr := make([]int, 0, n)
		sorted := make([]int, 0, n)
		for i := range n {
			a := nextInt()
			arr = append(arr, a)
			sorted = append(sorted, i+1)
		}

		ans := make([]int, n)
		for i, v := range arr {
			ans[i] = sorted[n-v]
		}
		for _, v := range ans {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1831/A
  Tags: greedy, implementation
  Rating: 800
*/

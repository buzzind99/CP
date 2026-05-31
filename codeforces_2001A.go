//go:build codeforces_2001A

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
		arr := make([]int, 101)
		maxCount := 0
		for range n {
			a := nextInt()
			arr[a]++
			maxCount = max(maxCount, arr[a])
		}

		fmt.Fprintln(wr, n-maxCount)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2001/A
  Tags: greedy, implementation
  Rating: 800
*/

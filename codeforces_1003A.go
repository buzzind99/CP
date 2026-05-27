//go:build codeforces_1003A

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

	n := nextInt()
	m := make(map[int]int)
	maxVal := -1
	for range n {
		a := nextInt()
		m[a]++
		maxVal = max(maxVal, m[a])
	}

	fmt.Fprintln(wr, maxVal)
}

/*
  Link: https://codeforces.com/problemset/problem/1003/A
  Tags: implementation
  Rating: 800
*/

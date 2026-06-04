//go:build codeforces_1968B

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
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, m := nextInt(), nextInt()
		a, b := next(), next()

		i, j := 0, 0
		for i < n && j < m {
			if a[i] == b[j] { i++ }
			j++
		}

		fmt.Fprintln(wr, i)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1968/B
  Tags: greedy, two pointers
  Rating: 800
*/

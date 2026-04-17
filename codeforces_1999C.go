//go:build codeforces_1999C

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
		n, s, m := nextInt(), nextInt(), nextInt()
		last := 0
		maxInterval := 0
		for range n {
			l, r := nextInt(), nextInt()
			maxInterval = max(maxInterval, l-last)
			last = r
		}
		maxInterval = max(maxInterval, m-last)

		if maxInterval < s {
			fmt.Fprintln(wr, "NO")
		} else {
			fmt.Fprintln(wr, "YES")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1999/C
  Tags: greedy, implementation
  Rating: 800
*/

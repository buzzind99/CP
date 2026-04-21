//go:build codeforces_2072A

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, k, p := nextInt(), nextInt(), nextInt()
		target := abs(k)
		ans := (target+p-1)/p

		if ans > n {
			fmt.Fprintln(wr, -1)
		} else {
			fmt.Fprintln(wr, ans)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2072/A
  Tags: greedy, implementation, math
  Rating: 800
*/

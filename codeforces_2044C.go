//go:build codeforces_2044C

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
		m, a, b, c := nextInt(), nextInt(), nextInt(), nextInt()
		remainderA, remainderB := max(0, m-a), max(0, m-b)
		cA := min(c, remainderA)
		cB := min(c-cA, remainderB)
		ans := (2*m - remainderA - remainderB) + cA + cB

		fmt.Fprintln(wr, ans)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2044/C
  Tags: greedy, math
  Rating: 800
*/

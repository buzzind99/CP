//go:build codeforces_1806A

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
		a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()

		if d < b || (a+d-b) < c {
			fmt.Fprintln(wr, -1)
		} else {
			ans := (d-b)+(a+d-b-c)
			fmt.Fprintln(wr, ans)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1806/A
  Tags: geometry, greedy, math
  Rating: 800
*/

//go:build codeforces_1547A

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
		xa, ya := nextInt(), nextInt()
		xb, yb := nextInt(), nextInt()
		xf, yf := nextInt(), nextInt()

		dist := abs(xa-xb)+abs(ya-yb)
		if xa == xb && xb == xf && yf > min(ya, yb) && yf < max(ya, yb) {
			dist += 2
		} else if ya == yb && yb == yf && xf > min(xa, xb) && xf < max(xa, xb) {
			dist += 2
		}

		fmt.Fprintln(wr, dist)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1547/A
  Tags: imlementation, math
  Rating: 800
*/

//go:build codeforces_1729A

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
		a, b, c := nextInt(), nextInt(), nextInt()
		aT := a-1
		bT := abs(c-b) + (c-1)
		sol := 1
		if c == 1 { bT = b-1 }
		if bT < aT { sol = 2 } else if bT == aT { sol = 3 }

		fmt.Fprintln(wr, sol)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1729/A
  Tags: math
  Rating: 800
*/

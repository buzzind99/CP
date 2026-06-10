//go:build codeforces_1829C

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

	const INF int = 1e9
	t := nextInt()
	for range t {
		n := nextInt()
		min11, min10, min01 := INF, INF, INF
		for range n {
			m, s := nextInt(), next()
			if s == "11" {
				min11 = min(min11, m)
			} else if s == "10" {
				min10 = min(min10, m)
			} else if s == "01" {
				min01 = min(min01, m)
			}
		}

		ansA := min11
		ansB := min10 + min01
		ans := min(ansA, ansB)

		if ans >= INF {
			fmt.Fprintln(wr, -1)
		} else {
			fmt.Fprintln(wr, ans)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1829/C
  Tags: bitmasks, greedy, implementation
  Rating: 800
*/

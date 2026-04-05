//go:build codeforces_1389A

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
		l, r := nextInt(), nextInt()

		if 2*l <= r {
			fmt.Fprintln(wr, l, 2*l)
		} else {
			fmt.Fprintln(wr, -1, -1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1389/A
  Tags: constructive algorithms, greedy, math, number theory
  Rating: 800
*/

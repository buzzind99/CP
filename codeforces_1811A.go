//go:build codeforces_1811A

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
		n, d, s := nextInt(), next(), next()
		inserted := false
		for i := 0; i < n; i++ {
			if !inserted && d[0] > s[i] {
				fmt.Fprint(wr, d)
				inserted = true
			}
			fmt.Fprint(wr, string(s[i]))
		}

		if !inserted {
			fmt.Fprint(wr, d)
		}
		fmt.Fprintln(wr)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1811/A
  Tags: greedy, math, strings
  Rating: 800
*/

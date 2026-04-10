//go:build codeforces_2074A

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
		l, r, d, u := nextInt(), nextInt(), nextInt(), nextInt()

		if l == r && l == r && l == d && l == u {
			fmt.Fprintln(wr, "Yes")
		} else {
			fmt.Fprintln(wr, "No")
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2074/A
  Tags: geometry, implementation
  Rating: 800
*/

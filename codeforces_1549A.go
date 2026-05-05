//go:build codeforces_1549A

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
		p := nextInt()

		fmt.Fprintln(wr, 2, p-1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1549/A
  Tags: math, number theory
  Rating: 800
*/

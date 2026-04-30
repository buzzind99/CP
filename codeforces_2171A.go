//go:build codeforces_2171A

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

func isEven (n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n := nextInt()

		if !isEven(n) {
			fmt.Fprintln(wr, 0)
		} else {
			fmt.Fprintln(wr, n/4+1)
		}
	}
}

/*
  Link: https://codeforces.com/problemset/problem/2171/A
  Tags: brute force, math
  Rating: 800
*/

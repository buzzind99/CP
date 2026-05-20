//go:build codeforces_1918A

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

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	t := nextInt()
	for range t {
		n, m := nextInt(), nextInt()

		fmt.Fprintln(wr, m/2*n)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1918/A
  Tags: constructive algorithms, greedy, implementation, math
  Rating: 800
*/

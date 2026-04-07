//go:build codeforces_1325A

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
		n := nextInt()

		fmt.Fprintln(wr, 1, n-1)
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1325/A
  Tags: constructive algorithms, greedy, number theory
  Rating: 800
*/

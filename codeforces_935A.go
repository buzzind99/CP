//go:build codeforces_935A

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

	n := nextInt()
	count := 0
	for leader := 1; leader <= n; leader++ {
		if n-leader < leader { break }
		if (n-leader)%leader == 0 { count++ }
	}

	fmt.Fprintln(wr, count)
}

/*
  Link: https://codeforces.com/problemset/problem/935/A
  Tags: brute force, implementation
  Rating: 800
*/

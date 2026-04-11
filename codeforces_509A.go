//go:build codeforces_509A

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
	sol := 1
	for i := 0; i < n-1; i++ {
		sol = sol*(n+i)/(i+1)
	}

	fmt.Fprintln(wr, sol)
}

/*
  Link: https://codeforces.com/problemset/problem/509/A
  Tags: brute force, implementation
  Rating: 800
*/

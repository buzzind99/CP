//go:build codeforces_822A

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

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	a, b := nextInt(), nextInt()

	fmt.Fprintln(wr, factorial(min(a, b)))
}

/*
  Link: https://codeforces.com/problemset/problem/822/A
  Tags: implementation, math, number theory
  Rating: 800
*/

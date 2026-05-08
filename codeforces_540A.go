//go:build codeforces_540A

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

func abs(x int) int {
	if x < 0 { return -x }
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	defer wr.Flush()

	n := nextInt()
	a, b := next(), next()
	sum := 0
	for i := range n {
		intA, intB := int(a[i]-'0'), int(b[i]-'0')
		low, high := min(intA, intB), max(intA, intB)
		left, right := 10-high+low, high-low
		sum += min(left, right)
	}

	fmt.Fprintln(wr, sum)
}

/*
  Link: https://codeforces.com/problemset/problem/540/A
  Tags: implementation
  Rating: 800
*/

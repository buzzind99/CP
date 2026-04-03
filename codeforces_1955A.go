//go:build codeforces_1955A

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
		n, a, b := nextInt(), nextInt(), nextInt()
		nDiscount := n/2
		withDiscount := nDiscount*b+(n-nDiscount*2)*a
		noDiscount := n*a
		fmt.Fprintln(wr, min(withDiscount, noDiscount))
	}
}

/*
  Link: https://codeforces.com/problemset/problem/1955/A
  Tags: math
  Rating: 800
*/
